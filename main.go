package main

import (
	"flag"
	"fmt"
	"os"
)

var version string = "0.0.1"

func exitWithError(errorCode int, errorMessage string) {
	print("[ERROR] %s\n", errorMessage)
	os.Exit(errorCode)
}

func print(message string, args ...interface{}) {
	fmt.Printf("[bugsnag-to-csv] "+message+"\r\n", args...)
}

func printHeader() {
	print("##################################################")
	print("#                                                #")
	print("#              Bugsnag-to-CSV, v%s            #", version)
	print("#                Xander Jones, 2021              #")
	print("# https://github.com/xander-jones/bugsnag-to-csv #")
	print("#                                                #")
	print("##################################################")
}

func main() {

	token := flag.String("token", "", "[REQUIRED][String] Your Bugsnag personal auth token.")
	getProjectIds := flag.Bool("show-project-ids", false, "Use this flag to get a list of project IDs accessible with your token.")
	//outputFilepath := flag.String("output-file", "", "[String] Filepath to store the downloaded CSV.")
	//projectId := flag.String("projectid", "", "The Project ID you wish to download from")
	//filters := flag.String("filters", "", "A JSON string array of filters to apply")
	flag.Parse()

	printHeader()

	if *token == "" {
		exitWithError(1, "Missing token. Please supply Bugsnag personal auth token with --token flag")
	}

	if *getProjectIds {
		print("getting your project IDs with token ")
	}

	// orgs := getUsersOrganizations(false, 30)
	// for _, org := range orgs {
	// 	fmt.Println(org["name"])
	// 	projects := getOrganizationsProjects(org["id"].(string), 10)
	// 	for _, proj := range projects {
	// 		fmt.Println(proj["name"])
	// 	}
	// }
}
