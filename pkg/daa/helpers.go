package daa

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func addQueryParams(url string) string {
	url += "?"
	if !common.MinimalReports {
		url += "full_reports=true"
	}
	// TODO: Add filters here. Perhaps they should go as the body though to save translating JSON to URL query.
	return url
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
