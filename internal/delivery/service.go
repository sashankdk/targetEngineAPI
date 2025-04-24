package delivery

import (
	"context"
	"errors"
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

var ErrAppId = errors.New("appId is required")
var ErrCountry = errors.New("country is required")
var ErrOs = errors.New("os is required")

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
