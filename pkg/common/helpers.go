package common

import (
	"fmt"
	"os"
)

var Verbose bool = false

func ExitWithError(errorCode int, errorMessage string) {
	Print("[ERROR] %s\n", errorMessage)
	os.Exit(errorCode)
}

func Print(message string, args ...interface{}) {
	fmt.Printf("[bugsnag-to-csv] "+message+"\r\n", args...)
}

func PrintVerbose(message string, args ...interface{}) {
	if Verbose {
		Print(message, args)
	}
}

func PrintHeader() {
	if Verbose {
		Print("##################################################")
		Print("#                                                #")
		Print("#              Bugsnag-to-CSV, v%s            #", PackageVersion)
		Print("#                Xander Jones, 2021              #")
		Print("# https://github.com/xander-jones/bugsnag-to-csv #")
		Print("#                                                #")
		Print("##################################################")
	}
}
