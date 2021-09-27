package daa

import "github.com/xander-jones/bugsnag-exporter/pkg/common"

func GetUsersOrganizations(admin bool, perPage int) []map[string]interface{} {
	// GET https://api.bugsnag.com/user/organizations
	var url string = "https://api.bugsnag.com/user/organizations"
	common.PrintVerbose("Getting user's organizations from API: " + url)
	var organizations []map[string]interface{} = BugsnagGetAllElements(url)
	return organizations
}
