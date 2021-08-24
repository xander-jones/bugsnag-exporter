package user

import "encoding/json"

func getUsersOrganizations(admin bool, perPage int) []map[string]interface{} {
	// GET https://api.bugsnag.com/user/organizations?admin=&per_page=10
	res := makeBugsnagDAAGet("https://api.bugsnag.com/user/organizations?admin=&per_page=10")

	var organizations []map[string]interface{}
	json.Unmarshal([]byte(res), &organizations)
	return organizations
}
