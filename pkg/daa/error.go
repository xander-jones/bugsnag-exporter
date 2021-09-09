package daa

import "encoding/json"

func GetErrorEvents(project_id string, error_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-an-error
	//   GET https://api.bugsnag.com/projects/project_id/errors/error_id/events
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/projects/" + project_id + "/errors/" + project_id + "/events")
	var events []map[string]interface{}
	json.Unmarshal([]byte(res), &events)
	return events
}
