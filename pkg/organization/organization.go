package organization

import "encoding/json"

func getOrganizationsProjects(organization_id string, per_page int) []map[string]interface{} {
	// GET https://api.bugsnag.com/organizations/organization_id/projects?q=MyProject&sort=created_at&direction=desc&per_page=30
	res := makeBugsnagDAAGet("https://api.bugsnag.com/organizations/" + organization_id + "/projects?per_page=10")
	var projects []map[string]interface{}
	json.Unmarshal([]byte(res), &projects)
	return projects
}
