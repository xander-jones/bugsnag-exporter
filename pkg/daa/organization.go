package daa

import (
	"encoding/json"
)

func GetOrganizationsProjects(organization_id string, per_page int) []map[string]interface{} {
	// GET https://api.bugsnag.com/organizations/organization_id/projects?q=MyProject&sort=created_at&direction=desc&per_page=30
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/organizations/" + organization_id + "/projects?per_page=10")
	var projects []map[string]interface{}
	json.Unmarshal([]byte(res), &projects)
	return projects
}
