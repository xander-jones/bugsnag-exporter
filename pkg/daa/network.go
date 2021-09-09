package daa

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BugsnagDAAResponse struct {
	body        []byte              // the JSON body response
	ratelimit   BugsnagDAARateLimit // The number of API calls that can be made per minute
	retry_after []string            // The datetime stamp of when a retry should be made (rate limit refresh)
	link        []string            // The link for the next set of data, if it exists
	totalcount  []string            // The total number of errors or event objects in this search
}
type BugsnagDAARateLimit struct {
	limit     []string
	remaining []string
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
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode == 200 {
		response.body = body
		response.ratelimit.limit = res.Header["X-RateLimit-Limit"]
		response.ratelimit.remaining = res.Header["X-RateLimit-Remaining"]
		response.retry_after = res.Header["Retry-After"]
		response.link = res.Header["Link"]
		response.totalcount = res.Header["X-Total-Count"]
		return response
	} else {
		fmt.Println("[HTTP Error] Non HTTP/200 response received: " + res.Status)
		response.body = body
		return response
	}
}
