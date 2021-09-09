package daa

import (
	"encoding/json"
)

func GetOrganizationsProjects(organization_id string, per_page int) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/projects/projects/list-an-organization's-projects
	//   GET https://api.bugsnag.com/organizations/organization_id/projects
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/organizations/" + organization_id + "/projects")
	var projects []map[string]interface{}
	json.Unmarshal([]byte(res.body), &projects)
	return projects
}
