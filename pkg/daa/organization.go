package daa

import "github.com/xander-jones/bugsnag-to-csv/pkg/common"

func GetOrganizationsProjects(organization_id string, per_page int) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/projects/projects/list-an-organization's-projects
	//   GET https://api.bugsnag.com/organizations/organization_id/projects
	var url string = "https://api.bugsnag.com/organizations/" + organization_id + "/projects"
	common.PrintVerbose("Getting organization projects from API: " + url)
	var projects []map[string]interface{} = BugsnagGetAllElements(url)
	return projects
}
