package repository
import (
	"targetApi/internal/model"
)

type Repository interface {
	GetActiveCampaigns() ([]model.Campaign, error)
	GetTargetingRules() (map[string]model.TargetingRule, error)
}


