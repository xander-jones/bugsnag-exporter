package daa

import (
	"encoding/json"

	"github.com/xander-jones/bugsnag-to-csv/pkg/common"
)

func GetError(project_id string, error_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/view-an-error
	//   GET https://api.bugsnag.com/projects/project_id/errors/error_id
	url := "https://api.bugsnag.com/projects/" + project_id + "/errors/" + error_id
	common.PrintVerbose("Getting error from API: " + url)
	res := MakeBugsnagDAAGet(url)
	PrintHeaders(res)

	var events []map[string]interface{}
	json.Unmarshal([]byte(res.body), &events)
	return events
}

func GetErrorEvents(project_id string, error_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-an-error
	//   GET https://api.bugsnag.com/projects/project_id/errors/error_id/events
	url := "https://api.bugsnag.com/projects/" + project_id + "/errors/" + error_id + "/events"
	common.PrintVerbose("Getting events from API: " + url)
	res := MakeBugsnagDAAGet(url)
	PrintHeaders(res)

	var events []map[string]interface{}
	json.Unmarshal([]byte(res.body), &events)
	return events
}
