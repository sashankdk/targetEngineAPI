package db

import (
	"context"
	"encoding/json"
	"fmt"

	"targetApi/internal/model"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(redisHost string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
	return &Cache{rdb: rdb}
}

func (c *Cache) SetCampaigns(ctx context.Context, campaigns []model.Campaign) error {
	data, err := json.Marshal(campaigns)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, "campaigns", data, 0).Err()
}

func (c *Cache) GetCampaigns(ctx context.Context) ([]model.Campaign, error) {
	val, err := c.rdb.Get(ctx, "campaigns").Result()
	if err != nil {
		return nil, err
	}

	var campaigns []model.Campaign
	err = json.Unmarshal([]byte(val), &campaigns)
	return campaigns, err
}

func (c *Cache) SetRules(ctx context.Context, rules map[string]model.TargetingRule) error {
	for cid, rule := range rules {
		data, err := json.Marshal(rule)
		if err != nil {
			return err
		}
		err = c.rdb.Set(ctx, fmt.Sprintf("rules:%s", cid), data, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) GetRule(ctx context.Context, cid string) (*model.TargetingRule, error) {
	val, err := c.rdb.Get(ctx, fmt.Sprintf("rules:%s", cid)).Result()
	if err != nil {
		return nil, err
	}

	var rule model.TargetingRule
	err = json.Unmarshal([]byte(val), &rule)
	return &rule, err
}
