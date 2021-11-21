package main

import (
	"flag"
	"fmt"

	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/xander-jones/bugsnag-exporter/pkg/common"
	"github.com/xander-jones/bugsnag-exporter/pkg/daa"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "575d0c15fae9fa9c865ede8258dea307",
		AppVersion:   common.PackageVersion,
		ReleaseStage: "development",
		ProjectPackages: []string{
			"main",
			"github.com/xander-jones/bugsnag-exporter",
			"github.com/xander-jones/bugsnag-exporter/common",
			"github.com/xander-jones/bugsnag-exporter/daa",
			"github.com/xander-jones/bugsnag-exporter/writers",
		},
		Synchronous: true,
	})

	token := flag.String("token", "", "[REQUIRED][String] Your Bugsnag personal auth token")
	getProjectIds := flag.Bool("show-project-ids", false, "[Flag] Use this flag to get a list of project IDs accessible with your token. Overrides any other flags")
	affectedUsers := flag.Bool("affected-users", false, "[Flag] Download a list of users affected by a specific error. Requires --project-id and --error-id")
	projectId := flag.String("project-id", "", "[String] The Project ID you wish to download from")
	errorId := flag.String("error-id", "", "[String] An error ID to download. If provided, downloads all events within filters for this error ID")
	events := flag.Bool("events", false, "[Flag] Download events rather than error groups when this flag is enabled. Requires --project-id (and optionally --error-id)")
	outputDir := flag.String("output-dir", "./data", "[String] Directory to store the downloaded file")
	filters := flag.String("filters", "", "[String] A string array of filters to apply (URL format)")
	//rateLimit := flag.Int("rate-limit", 0, "[Int] Set the number of calls to make per minute. Defaults to 0, no rate limit")
	minimalReports := flag.Bool("minimal", false, "[Flag] Download minimal event reports only for smaller file sizes")
	useCsv := flag.Bool("csv", false, "[Flag] Output data to file as CSV. Default false, noramally outputs as JSON")
	noWarn := flag.Bool("no-warn", false, "[Flag] Don't warn me if this call will take more than 5 calls to the API")
	verbose := flag.Bool("verbose", false, "[Flag] Set the output to be verbose for debugging purposes.")
	flag.Parse()

	common.Verbose = *verbose
	common.NoWarn = *noWarn
	common.UseCsv = *useCsv
	common.OutputDir = *outputDir
	common.MinimalReports = *minimalReports

	if *token == "" {
		common.ExitWithString(1, "Missing token. Please supply Bugsnag personal auth token with --token flag")
	} else {
		daa.PersonalAuthToken = *token
		if *getProjectIds {
			common.Print("Getting your project IDs with provided token")
			orgs := daa.GetUsersOrganizations()
			for _, org := range orgs {
				common.Print("Organization: " + fmt.Sprint(org["name"]) + " [" + fmt.Sprint(org["id"]) + "]")
				projects := daa.GetOrganizationsProjects(org["id"].(string))
				for _, proj := range projects {
					common.Print("  > " + fmt.Sprint(proj["name"]) + " [" + fmt.Sprint(proj["id"]) + "]")
				}
			}
		} else if *affectedUsers {
			if *errorId != "" && *projectId != "" {
				common.Print("Downloading list of users affected by errorId '%s' on projectId '%s'", *errorId, *projectId)
				daa.GetAffectedUsers(*projectId, *errorId, *filters)
			} else {
				common.ExitWithString(1, "--affected-users requires --project-id and --error-id to get list of affected users")
			}
		} else {
			if *projectId == "" {
				common.ExitWithString(2, "Missing Project ID. Please supply a Project ID with --project-id flag")
			} else {
				if *errorId == "" {
					if *events {
						common.Print("Downloading all events for projectId within filters")
						daa.GetProjectEvents(*projectId, *filters)
					} else {
						common.Print("Downloading all errors from projectId within filters")
						daa.GetProjectErrors(*projectId, *filters)
					}
				} else {
					if *events {
						common.Print("Downloading all events for projectId & errorId within filters")
						daa.GetErrorEvents(*projectId, *errorId, *filters)
					} else {
						common.Print("Downloading error from projectId & errorId within filters")
						daa.GetError(*projectId, *errorId, *filters)
					}
				}
			}
		}
	}
}
