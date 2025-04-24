package middleware

import (
	"targetApi/internal/delivery"
	"targetApi/internal/model"

	"github.com/sony/gobreaker"
)

func CircuitBreakerMiddleware(cb *gobreaker.CircuitBreaker) func(delivery.DeliveryService) delivery.DeliveryService {
	return func(next delivery.DeliveryService) delivery.DeliveryService {
		return &cbService{cb, next}
	}
}

type cbService struct {
	cb   *gobreaker.CircuitBreaker
	next delivery.DeliveryService
}

func (s *cbService) GetCampaigns(appID, country, os string) ([]model.Campaign, error) {
	result, err := s.cb.Execute(func() (interface{}, error) {
		return s.next.GetCampaigns(appID, country, os)
	})
	if err != nil {
		return nil, err
	}
	return result.([]model.Campaign), nil
}
