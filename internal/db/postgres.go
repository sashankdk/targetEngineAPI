package db

import (
	"database/sql"
	"targetApi/internal/model"
)

type PgRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PgRepo {
	return &PgRepo{db: db}
}

func (r *PgRepo) GetActiveCampaigns() ([]model.Campaign, error) {
	rows, err := r.db.Query(`SELECT id, img, cta FROM campaigns WHERE status = 'ACTIVE'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Campaign
	for rows.Next() {
		var c model.Campaign
		if err := rows.Scan(&c.ID, &c.Img, &c.CTA); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (r *PgRepo) GetTargetingRules() (map[string]model.TargetingRule, error) {
	rows, err := r.db.Query(`
		SELECT campaign_id, COALESCE(include_app, '{}'::text[]), COALESCE(exclude_app, '{}'::text[]), COALESCE(include_os, '{}'::text[]), COALESCE(exclude_os, '{}'::text[]), COALESCE(include_country, '{}'::text[]), COALESCE(exclude_country, '{}'::text[])
		FROM targeting_rules`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]model.TargetingRule)
	for rows.Next() {
		var tr model.TargetingRule
		err := rows.Scan(
			&tr.CampaignID,
			&tr.IncludeApp, &tr.ExcludeApp,
			&tr.IncludeOS, &tr.ExcludeOS,
			&tr.IncludeCountry, &tr.ExcludeCountry,
		)
		if err != nil {
			return nil, err
		}
		result[tr.CampaignID] = tr
	}
	return result, nil
}
