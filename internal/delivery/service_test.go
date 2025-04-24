package delivery

import (
	"targetApi/internal/model"
	"testing"
)

func TestMatches(t *testing.T) {
	tests := []struct {
		name    string
		rule    *model.TargetingRule
		appID   string
		country string
		os      string
		want    bool
	}{
		{
			name: "Include all valid",
			rule: &model.TargetingRule{
				IncludeApp:     []string{"app1"},
				IncludeCountry: []string{"IN"},
				IncludeOS:      []string{"android"},
			},
			appID: "app1", country: "IN", os: "android", want: true,
		},
		{
			name: "Excluded country",
			rule: &model.TargetingRule{
				ExcludeCountry: []string{"US"},
			},
			appID: "app2", country: "US", os: "ios", want: false,
		},
		{
			name: "Included OS but not matched",
			rule: &model.TargetingRule{
				IncludeOS: []string{"android"},
			},
			appID: "app3", country: "FR", os: "ios", want: false,
		},
		{
			name:  "No rules (default allow)",
			rule:  &model.TargetingRule{},
			appID: "xyz", country: "BR", os: "web", want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matches(tt.rule, tt.appID, tt.country, tt.os)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
