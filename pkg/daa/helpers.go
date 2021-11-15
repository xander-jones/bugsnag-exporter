package daa

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func addQueryParams(url string, filters string) string {
	url += "?"
	if !common.MinimalReports {
		url += "full_reports=true"
	}
	if filters != "" {
		if f := formatFilters(filters); f != "BAD_FILTERS" {
			url += "&"
			url += f
		} else {
			common.ExitWithString(1, "Filters provided were not valid")
		}
	}
	return url
}

func formatFilters(filters string) string {
	r := regexp.MustCompile(`\??(filters\[[^]]*\]=[^&]*)&?`)
	matches := r.FindAllString(filters, -1)
	common.PrintVerbose("filters match count: " + fmt.Sprint(len(matches)))
	common.PrintVerbose("filters: " + fmt.Sprint(matches))
	if matches == nil {
		return "BAD_FILTERS"
	} else {
		return filters
	}
}

/*
	Waits for a user to confirm a message. If they reject the message, the
	application will exit with a message.

	With thanks to @r0l1: https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
*/
func getConfirmation(confirmationMessage string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", confirmationMessage)
		response, err := reader.ReadString('\n')
		if err != nil {
			common.ExitWithError(1, err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			common.ExitWithString(1, "User rejected further API calls")
		}
	}
}
