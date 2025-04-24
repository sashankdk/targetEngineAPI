package listener

import (
	"context"
	"log"
	"time"

	"targetApi/internal/db"

	"github.com/lib/pq"
)

func ListenForCampaignChanges(dbConn string, pg *db.PgRepo, cache *db.Cache) {
	listener := pq.NewListener(dbConn, 1*time.Second, 10*time.Second, nil)

	// Listen on multiple channels
	channels := []string{"campaign_change", "targeting_rule_change"}
	for _, channel := range channels {
		if err := listener.Listen(channel); err != nil {
			log.Fatalf("Failed to listen on %s: %v", channel, err)
		}
	}

	log.Println("Listening for campaign/target-rule changes changes...")

	for {
		select {
		case notification := <-listener.Notify:
			if notification == nil {
				continue
			}

			ctx := context.Background()
			log.Printf("Change detected on channel: %s, payload: %s", notification.Channel, notification.Extra)

			switch notification.Channel {
			case "campaign_change":
				if err := refreshCampaignCache(ctx, pg, cache); err != nil {
					log.Printf("Error refreshing campaigns: %v", err)
				}
			case "targeting_rule_change":
				if err := refreshTargetingRulesCache(ctx, pg, cache); err != nil {
					log.Printf("Error refreshing targeting rules: %v", err)
				}
			default:
				log.Printf("Unknown notification channel: %s", notification.Channel)
			}

		case <-time.After(90 * time.Second):
			log.Println("⏱ Re-pinging Postgres listener...")
			go listener.Ping()
		}
	}
}

func refreshCampaignCache(ctx context.Context, pg *db.PgRepo, cache *db.Cache) error {
	campaigns, err := pg.GetActiveCampaigns()
	if err != nil {
		return err
	}
	return cache.SetCampaigns(ctx, campaigns)
}

func refreshTargetingRulesCache(ctx context.Context, pg *db.PgRepo, cache *db.Cache) error {
	rules, err := pg.GetTargetingRules()
	if err != nil {
		return err
	}
	return cache.SetRules(ctx, rules)
}
