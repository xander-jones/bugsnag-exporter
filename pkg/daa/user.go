package daa

import "github.com/xander-jones/bugsnag-exporter/pkg/common"

func GetUsersOrganizations() []map[string]interface{} {
	// GET https://api.bugsnag.com/user/organizations
	var url string = "https://api.bugsnag.com/user/organizations"
	common.PrintVerbose("Getting user's organizations from API: %s", url)
	var organizations []map[string]interface{} = BugsnagGetArray(url)
	return organizations
}
