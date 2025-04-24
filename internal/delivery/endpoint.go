package delivery

import (
	"context"
	"targetApi/internal/model"

	"github.com/go-kit/kit/endpoint"
)

type GetdeliveryRequest struct {
	AppId   string `json:"appId"`
	Country string `json:"country"`
	Os      string `json:"os"`
}

type GetdeliveryResponse struct {
	Campaigns []model.Campaign `json:"-"`
	Err       string           `json:"error,omitempty"`
}

func GetdeliveryEndpoint(svc DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetdeliveryRequest)

		if req.AppId == "" || req.Country == "" || req.Os == "" {
			return GetdeliveryResponse{Err: "missing app/country/os param"}, nil
		}
		campaigns, err := svc.GetCampaigns(req.AppId, req.Country, req.Os)
		if err != nil {
			return GetdeliveryResponse{Err: err.Error()}, nil
		}
		return GetdeliveryResponse{Campaigns: campaigns}, nil
	}
}
