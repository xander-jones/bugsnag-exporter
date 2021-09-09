package daa

import (
	"encoding/json"
)

func GetUsersOrganizations(admin bool, perPage int) []map[string]interface{} {
	// GET https://api.bugsnag.com/user/organizations?admin=&per_page=10
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/user/organizations?admin=&per_page=10")
	var organizations []map[string]interface{}
	json.Unmarshal([]byte(res.body), &organizations)
	return organizations
}
