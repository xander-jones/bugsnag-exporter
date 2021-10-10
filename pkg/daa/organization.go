package daa

import "github.com/xander-jones/bugsnag-exporter/pkg/common"

func GetOrganizationsProjects(organizationId string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/projects/projects/list-an-organization's-projects
	//   GET https://api.bugsnag.com/organizations/organization_id/projects
	var url string = "https://api.bugsnag.com/organizations/" + organizationId + "/projects"
	common.PrintVerbose("Getting organization projects from API: " + url)
	var projects []map[string]interface{} = BugsnagGetArray(url)
	return projects
}
