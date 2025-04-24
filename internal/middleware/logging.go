package middleware

import (
	"targetApi/internal/delivery"
	"targetApi/internal/model"
	"time"

	"github.com/go-kit/log"
)

func LoggingMiddleware(logger log.Logger) func(delivery.DeliveryService) delivery.DeliveryService {
	return func(next delivery.DeliveryService) delivery.DeliveryService {
		return &loggedService{logger, next}
	}
}

type loggedService struct {
	logger log.Logger
	next   delivery.DeliveryService
}

func (s *loggedService) GetCampaigns(appID, country, os string) ([]model.Campaign, error) {
	start := time.Now()
	campaigns, err := s.next.GetCampaigns(appID, country, os)
	duration := time.Since(start)

	_ = s.logger.Log(
		"ts", time.Now().Format(time.RFC3339),
		"endpoint", "/v1/delivery",
		"app", appID,
		"country", country,
		"os", os,
		"status", statusCode(len(campaigns), err),
		"campaign_count", len(campaigns),
		"duration_ms", duration.Milliseconds(),
		"error", err,
	)

	return campaigns, err
}

func statusCode(count int, err error) int {
	if err != nil {
		return 500
	}
	if count == 0 {
		return 204
	}
	return 200
}
