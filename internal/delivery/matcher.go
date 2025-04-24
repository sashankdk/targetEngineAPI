package delivery
import (
	"targetApi/internal/model"
)
func matches(rule *model.TargetingRule, appID, country, os string) bool {
	// App ID check
	if len(rule.IncludeApp) > 0 && !contains(rule.IncludeApp, appID) {
		return false
	}
	if len(rule.ExcludeApp) > 0 && contains(rule.ExcludeApp, appID) {
		return false
	}

	// Country check
	if len(rule.IncludeCountry) > 0 && !contains(rule.IncludeCountry, country) {
		return false
	}
	if len(rule.ExcludeCountry) > 0 && contains(rule.ExcludeCountry, country) {
		return false
	}

	// OS check
	if len(rule.IncludeOS) > 0 && !contains(rule.IncludeOS, os) {
		return false
	}
	if len(rule.ExcludeOS) > 0 && contains(rule.ExcludeOS, os) {
		return false
	}

	// If all checks passed
	return true
}

func contains(list []string, target string) bool {
	for _, val := range list {
		if val == target {
			return true
		}
	}
	return false
}
