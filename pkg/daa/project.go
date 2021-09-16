package daa

import (
	"encoding/json"

	"github.com/xander-jones/bugsnag-to-csv/pkg/common"
)

func GetProjectErrors(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/list-the-errors-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/errors
	url := "https://api.bugsnag.com/projects/" + project_id + "/errors"
	common.PrintVerbose("Getting errors from API: " + url)
	res := MakeBugsnagDAAGet(url)

	var errors []map[string]interface{}
	json.Unmarshal([]byte(res.body), &errors)
	return errors
}

func GetProjectEvents(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/events
	url := "https://api.bugsnag.com/projects/" + project_id + "/events"
	common.PrintVerbose("Getting events from API: " + url)
	res := MakeBugsnagDAAGet(url)

	var events []map[string]interface{}
	json.Unmarshal([]byte(res.body), &events)
	return events
}