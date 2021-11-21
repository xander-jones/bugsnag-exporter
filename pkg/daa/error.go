package daa

import (
	"os"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
	"github.com/xander-jones/bugsnag-exporter/pkg/writers"
)

func GetError(projectId string, errorId string, filters string) map[string]interface{} {
	// Docs https://bugsnagapiv2.docs.apiary.io/#reference/errors/errors/view-an-error
	//  GET https://api.bugsnag.com/projects/project_id/errors/error_id
	var url string = addQueryParams("https://api.bugsnag.com/projects/"+projectId+"/errors/"+errorId, filters)
	common.PrintVerbose("Getting error from API: %s", url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "errorId-"+errorId)
	var err map[string]interface{} = BugsnagGetObject(url)
	writers.WriteObjectToFile(handle, err)
	writers.CloseOutputFile(handle)
	common.Print("Downloaded error. Saved to '%s'", handle.Name())
	return err
}

func GetErrorEvents(projectId string, errorId string, filters string) []map[string]interface{} {
	// Docs https://bugsnagapiv2.docs.apiary.io/#reference/errors/events/list-the-events-on-an-error
	//  GET https://api.bugsnag.com/projects/project_id/errors/error_id/events
	var url string = addQueryParams("https://api.bugsnag.com/projects/"+projectId+"/errors/"+errorId+"/events", filters)
	common.PrintVerbose("Getting events from API: %s", url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "events-of-errorId-"+errorId)
	var events []map[string]interface{} = BugsnagGetArray(url)
	writers.WriteArrayToFile(handle, events)
	writers.CloseOutputFile(handle)
	common.Print("Downloaded %d events from error. Saved to '%s'", len(events), handle.Name())
	return events
}

func GetAffectedUsers(projectId string, errorId string, filters string) []map[string]interface{} {
	// Docs https://bugsnagapiv2.docs.apiary.io/#reference/errors/pivots/list-values-of-a-pivot-on-an-error
	//  GET https://api.bugsnag.com/projects/project_id/errors/error_id/pivots/event_field_display_id/values
	var url string = addQueryParams("https://api.bugsnag.com/projects/"+projectId+"/errors/"+errorId+"/pivots/user.id/values", filters)
	common.PrintVerbose("Getting affected users from API: %s", url)
	var handle *os.File = writers.CreateNewOutputFile(projectId, "users-affected-by-errorId-"+errorId)
	var events []map[string]interface{} = BugsnagGetArray(url)
	writers.WriteArrayToFile(handle, events)
	writers.CloseOutputFile(handle)
	common.Print("Downloaded information about %d users affected by errorId '%s'. Saved to '%s'", len(events), errorId, handle.Name())
	return events
}
