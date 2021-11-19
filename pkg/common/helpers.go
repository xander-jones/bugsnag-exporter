package common

import (
	"fmt"
	"os"
)

/*
	Variable to control output verbosity.
	Defaults to false
*/
var Verbose bool = false

/*
	Variable to control whether a warning is given
	if more than 10 API calls will be required to complete
	an operation. 10 API calls would encounter rate limiting
	under normal conditions.
*/
var NoWarn bool = false

/*
	Variable to control whether we output as a JSON file
	or as a CSV. JSON is preferable as data in Bugsnag events is
	complex and does not suit CSVs well.
*/
var UseCsv bool = false

/*
	Variable to store the output directory to store the files created
*/
var OutputDir string = "./"

/*
	Variable to control whether we should be fetching minimal reports
*/
var MinimalReports bool = false

/*
	Exits the application with the error code given after
	printing an error message string to the console.
*/
func ExitWithString(errorCode int, errorMessage string) {
	Print("[ERROR] %s\n", errorMessage)
	os.Exit(errorCode)
}

/*
	Exits the application with the error code given after
	printing an error message from an error object to console.
*/
func ExitWithError(errorCode int, err error) {
	Print("[ERROR] %s\n", err.Error())
	os.Exit(errorCode)
}

/*
	Exits the application with the error code given after
	printing a string and message from an error object to console.
*/
func ExitWithErrorAndString(errorCode int, err error, errorMessage string) {
	Print("[ERROR] %s\n", errorMessage)
	Print("[ERROR] %s\n", err.Error())
	os.Exit(errorCode)
}

/*
	Prints a message to the console, can take in arguments
	after the initial message so you can treat this as a Sprintf()
*/
func Print(format string, args ...interface{}) {
	fmt.Printf("[bugsnag-exporter] "+format+"\r\n", args...)
}

/*
	Prints a message to the console, in the same way as Print()
	but ONLY if `Verbose` is set to true
*/
func PrintVerbose(format string, args ...interface{}) {
	if Verbose {
		fmt.Printf("[b-e-verbose-----] "+format+"\r\n", args...)
	}
}

/*
	Prints a header to the console, including the application version
*/
func PrintHeader() {
	if Verbose {
		Print("####################################################")
		Print("#                                                  #")
		Print("#              bugsnag-exporter, v%s            #", PackageVersion)
		Print("#                Xander Jones, 2021                #")
		Print("# https://github.com/xander-jones/bugsnag-exporter #")
		Print("#                                                  #")
		Print("####################################################")
	}
}
