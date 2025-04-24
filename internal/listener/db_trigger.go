package listener

import (
	"context"
	"log"
	"time"

	"targetApi/internal/db"

	"github.com/lib/pq"
)

func ListenForCampaignChanges(dbConn string, repo *db.PgRepo, cache *db.Cache) {
	listener := pq.NewListener(dbConn, 1*time.Second, 10*time.Second, nil)

	if err := listener.Listen("campaign_change"); err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	log.Println("📡 Listening for campaign changes...")

	for {
		select {
		case n := <-listener.Notify:
			log.Printf("🔄 Campaign change detected: %s. Refreshing cache...", n.Extra)
			ctx := context.Background()
			campaigns, err := repo.GetActiveCampaigns()
			if err == nil {
				cache.SetCampaigns(ctx, campaigns)
			}
			rules, err := repo.GetTargetingRules()
			if err == nil {
				cache.SetRules(ctx, rules)
			}

		case <-time.After(90 * time.Second):
			log.Println("⏱ Re-pinging Postgres...")
			go listener.Ping()
		}
	}
}
