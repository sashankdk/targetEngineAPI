package model

type Campaign struct {
	ID  string `json:"cid"`
	Img string `json:"img"`
	CTA string `json:"cta"`
}

type TargetingRule struct {
	CampaignID     string
	IncludeApp     []string
	ExcludeApp     []string
	IncludeOS      []string
	ExcludeOS      []string
	IncludeCountry []string
	ExcludeCountry []string
}
