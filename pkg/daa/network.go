package daa

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/xander-jones/bugsnag-to-csv/pkg/common"
)

type BugsnagDAAResponse struct {
	body        []byte              // the JSON body response
	rateLimit   BugsnagDAARateLimit // The number of API calls that can be made per minute
	retryAfter  string              // The datetime stamp of when a retry should be made (rate limit refresh)
	link        BugsnagDAANextLink  // The link for the next set of data, if it exists
	xTotalCount int64               // The total number of errors or event objects in this search
}
type BugsnagDAARateLimit struct {
	limit     int64
	remaining int64
}

type BugsnagDAANextLink struct {
	url string // The URL of the link in the `Link` header
	rel string // The direction of the link. Can be `next` or `prev`
}

func PrintHeaders(res BugsnagDAAResponse) {
	common.PrintVerbose("X-Total-Count: " + fmt.Sprint(res.xTotalCount))
	common.PrintVerbose("Ratelimit:     " + fmt.Sprint(res.rateLimit.limit))
	common.PrintVerbose("Remaining:     " + fmt.Sprint(res.rateLimit.remaining))
	common.PrintVerbose("Link:          " + fmt.Sprint(res.link))
	common.PrintVerbose("Retry-After:   " + fmt.Sprint(res.retryAfter))
}

var client = &http.Client{}
var PersonalAuthToken = "" // Personal Auth Token "Go-API"

func MakeBugsnagDAAGet(url string) BugsnagDAAResponse {
	var response BugsnagDAAResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "token "+PersonalAuthToken)

	res, err := client.Do(req)
	if err != nil {
		common.ExitWithError(1000, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		common.ExitWithError(1000, err)
	}
	// for key, values := range res.Header {
	// 	fmt.Printf("%s: %v\n", key, values)
	// }
	if res.StatusCode == 200 {
		common.PrintVerbose("[HTTP Success] HTTP/200 response received")
	} else {
		common.PrintVerbose("[HTTP Error] Non HTTP/200 response received: " + res.Status)
	}
	parseHeaderInt(res.Header["X-Ratelimit-Limit"], 1001, &response.rateLimit.limit)
	parseHeaderInt(res.Header["X-Ratelimit-Remaining"], 1002, &response.rateLimit.remaining)
	parseHeaderInt(res.Header["X-Total-Count"], 1003, &response.xTotalCount)

	response.retryAfter = canonicalHeader(res.Header["Retry-After"])
	response.link = parseNextHeader(res.Header["Link"])
	response.body = body

	return response
}

// Parse a header string number into an int64, throwing an error
// if an integer conversion does not succeed.
func parseHeaderInt(headerValuesArray []string, errorCode int, outputPtr *int64) {
	var err error
	*outputPtr, err = strconv.ParseInt(canonicalHeader(headerValuesArray), 10, 64)
	if err != nil {
		common.ExitWithErrorAndString(errorCode, err, "An API response header returned an unexpected non-integer value")
	}
}

/*
	Returns the first element in the headers array. Any more than one header for the ones
	we care about would be unexpected, so just return the first.
	If the array is empty, just return an empty string
*/
func canonicalHeader(headerValuesArray []string) string {
	if len(headerValuesArray) > 0 {
		return headerValuesArray[0]
	} else {
		return ""
	}
}

/*
	Extract the next link and which direction it goes in relation to the current URL
*/
func parseNextHeader(headerValuesArray []string) BugsnagDAANextLink {
	header := canonicalHeader(headerValuesArray)
	var rtn BugsnagDAANextLink
	if len(header) > 0 {
		r, _ := regexp.Compile("<(http[s]://api.bugsnag.com/.*)>; rel=\"(next|prev)\"")
		matches := r.FindAllStringSubmatch(header, -1)
		rtn.url = matches[0][1]
		rtn.rel = matches[0][2]
	}
	return rtn
}
