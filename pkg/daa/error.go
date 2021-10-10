package daa

import (
	"os"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
	"github.com/xander-jones/bugsnag-exporter/pkg/writers"
)

func GetError(projectId string, errorId string) map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/view-an-error
	//   GET https://api.bugsnag.com/projects/project_id/errors/error_id
	var url string = addQueryParams("https://api.bugsnag.com/projects/" + projectId + "/errors/" + errorId)
	common.PrintVerbose("Getting error from API: " + url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "error-"+errorId)
	var err map[string]interface{} = BugsnagGetObject(url)
	writers.WriteObjectToFile(handle, err)
	writers.CloseOutputFile(handle)
	return err
}

func GetErrorEvents(projectId string, errorId string) []map[string]interface{} {
	// Docs: https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-an-error
	//   GET https://api.bugsnag.com/projects/project_id/errors/error_id/events
	var url string = addQueryParams("https://api.bugsnag.com/projects/" + projectId + "/errors/" + errorId + "/events")
	common.PrintVerbose("Getting events from API: " + url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "events-of-"+errorId)
	var events []map[string]interface{} = BugsnagGetArray(url)
	writers.WriteArrayToFile(handle, events)
	writers.CloseOutputFile(handle)
	return events
}

func GetUsersAffected(project_id string, error_id string) []map[string]interface{} {
	// TODO: Get a list of users affected by an error
	return nil
}
