package model

import (
	"github.com/lib/pq"
)

type Campaign struct {
	ID  string `json:"cid"`
	Img string `json:"img"`
	CTA string `json:"cta"`
}

type TargetingRule struct {
	CampaignID     string
	IncludeApp     pq.StringArray
	ExcludeApp     pq.StringArray
	IncludeOS      pq.StringArray
	ExcludeOS      pq.StringArray
	IncludeCountry pq.StringArray
	ExcludeCountry pq.StringArray
}
