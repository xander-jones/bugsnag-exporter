package user

import (
	"encoding/json"

	"github.com/xander-jones/bugsnag-to-csv/pkg/common"
)

func GetUsersOrganizations(admin bool, perPage int) []map[string]interface{} {
	// GET https://api.bugsnag.com/user/organizations?admin=&per_page=10
	res := common.MakeBugsnagDAAGet("https://api.bugsnag.com/user/organizations?admin=&per_page=10")

	var organizations []map[string]interface{}
	json.Unmarshal([]byte(res), &organizations)
	return organizations
}
