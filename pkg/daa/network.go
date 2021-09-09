package daa

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var client = &http.Client{}
var PersonalAuthToken = "" // Personal Auth Token "Go-API"

func MakeBugsnagDAAGet(url string) []byte {
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
		return body
	} else {
		fmt.Println("[HTTP Error] Non HTTP/200 response received: " + res.Status)
		return body
	}
}
