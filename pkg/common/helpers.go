package common

import (
	"fmt"
	"os"
)

func ExitWithError(errorCode int, errorMessage string) {
	Print("[ERROR] %s\n", errorMessage)
	os.Exit(errorCode)
}

func Print(message string, args ...interface{}) {
	fmt.Printf("[bugsnag-to-csv] "+message+"\r\n", args...)
}

func PrintVerbose(verbose *bool, message string, args ...interface{}) {
	if *verbose {
		Print(message, args)
	}
}

func PrintHeader(verbose *bool) {
	if *verbose {
		Print("##################################################")
		Print("#                                                #")
		Print("#              Bugsnag-to-CSV, v%s            #", PackageVersion)
		Print("#                Xander Jones, 2021              #")
		Print("# https://github.com/xander-jones/bugsnag-to-csv #")
		Print("#                                                #")
		Print("##################################################")
	}
}
