package daa

import (
	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func GetProjectErrors(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/list-the-errors-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/errors
	var url string = addQueryParams("https://api.bugsnag.com/projects/" + project_id + "/errors")
	common.PrintVerbose("Getting errors from API: " + url)
	var errs []map[string]interface{} = BugsnagGetArray(url)
	return errs
}

func GetProjectEvents(project_id string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-a-project
	//   GET https://api.bugsnag.com/projects/project_id/events
	var url string = addQueryParams("https://api.bugsnag.com/projects/" + project_id + "/events")
	common.PrintVerbose("Getting events from API: " + url)
	var events []map[string]interface{} = BugsnagGetArray(url)
	return events
}
