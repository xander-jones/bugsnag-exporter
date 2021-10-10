package daa

import "github.com/xander-jones/bugsnag-exporter/pkg/common"

func addQueryParams(url string) string {
	url += "?"
	if !common.MinimalReports {
		url += "full_reports=true"
	}
	// TODO: Add filters here. Perhaps they should go as the body though to save translating JSON to URL query.
	return url
}
