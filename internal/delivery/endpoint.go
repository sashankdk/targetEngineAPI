package delivery

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type GetdeliveryRequest struct{
	AppId   string `json:"appId"`
	Country string `json:"country"`
	Os      string `json:"os"`
}

type GetdeliveryResponse struct {
	Campaigns []Campaign `json:"campaigns,omitempty"`
	Err       string    `json:"err,omitempty"`
}

func GetdeliveryEndpoint(svc DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetdeliveryRequest)
		campaigns, err := svc.GetCampaigns(req.AppId, req.Country, req.Os)
		if err != nil {
			return GetdeliveryResponse{Err: err.Error()}, nil
		}
		return GetdeliveryResponse{Campaigns: campaigns}, nil
	}
}