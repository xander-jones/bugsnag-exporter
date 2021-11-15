package daa

import (
	"os"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
	"github.com/xander-jones/bugsnag-exporter/pkg/writers"
)

func GetProjectErrors(projectId string, filters string) []map[string]interface{} {
	// Docs https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/list-the-errors-on-a-project
	//  GET https://api.bugsnag.com/projects/project_id/errors
	var url string = addQueryParams("https://api.bugsnag.com/projects/"+projectId+"/errors", filters)
	common.PrintVerbose("Getting errors from API: " + url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "errors-of-projectId-"+projectId)
	var errs []map[string]interface{} = BugsnagGetArray(url)
	writers.WriteArrayToFile(handle, errs)
	writers.CloseOutputFile(handle)
	return errs
}

func GetProjectEvents(projectId string, filters string) []map[string]interface{} {
	// Docs https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-a-project
	//  GET https://api.bugsnag.com/projects/project_id/events
	var url string = addQueryParams("https://api.bugsnag.com/projects/"+projectId+"/events", filters)
	common.PrintVerbose("Getting events from API: " + url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "events-of-projectId-"+projectId)
	var events []map[string]interface{} = BugsnagGetArray(url)
	writers.WriteArrayToFile(handle, events)
	writers.CloseOutputFile(handle)
	return events
}
