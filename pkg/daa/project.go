package daa

import "encoding/json"

func GetProjectErrors(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/list-the-errors-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/errors
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/projects/" + project_id + "/errors")
	var errors []map[string]interface{}
	json.Unmarshal([]byte(res), &errors)
	return errors
}

func GetProjectEvents(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/events
	res := MakeBugsnagDAAGet("https://api.bugsnag.com/projects/" + project_id + "/events")
	var events []map[string]interface{}
	json.Unmarshal([]byte(res), &events)
	return events
}
