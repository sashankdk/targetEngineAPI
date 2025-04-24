package delivery

import (
	"context"
	"targetApi/internal/db"
	"targetApi/internal/model"
)

type DeliveryService interface {
	GetCampaigns(appId, country, os string) ([]model.Campaign, error)
}

type deliveryStruct struct {
	cache *db.Cache
}

func NewService(cache *db.Cache) DeliveryService {
	return &deliveryStruct{cache: cache}
}

func (d *deliveryStruct) GetCampaigns(appId, country, os string) ([]model.Campaign, error) {

	ctx := context.Background()
	campaigns, err := d.cache.GetCampaigns(ctx)
	if err != nil {
		return nil, err
	}
	var result []model.Campaign
	for _, c := range campaigns {
		rule, err := d.cache.GetRule(ctx, c.ID)
		if err != nil {
			continue
		}
		if matches(rule, appId, country, os) {
			result = append(result, c)
		}
	}
	return result, nil
}
