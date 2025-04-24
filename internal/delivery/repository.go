package delivery

import (
	"errors"
)


type Campaign struct {
	ID          string `json:"cid"`
	Img 	   string `json:"img"`
	CTA 	   string `json:"cta"`
}

type DeliveryService  interface {
	GetCampaigns(appId, country , os string) ([]Campaign, error)
}


type deliveryStruct struct {}

func NewService() DeliveryService {
	return &deliveryStruct{}
}
var ErrAppId = errors.New("appId is required")
var ErrCountry = errors.New("country is required")
var ErrOs = errors.New("os is required")

func (d *deliveryStruct) GetCampaigns(appId, country, os string) ([]Campaign, error) {

	if appId == "" {
		return nil, ErrAppId
	}
	if country == "" {
		return nil, ErrCountry
	}
	if os == "" {
		return nil, ErrOs
	}

	// Integrate with redis later

	return []Campaign{
		{
			ID:     "1",
			Img:    "https://example.com/image1.jpg",
			CTA:    "Click here",
		}},nil
}