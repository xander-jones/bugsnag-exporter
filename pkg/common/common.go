package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var client = &http.Client{}

const DAA_TOKEN = "e65783aa-bb55-4176-9080-19c101cf650f" // Personal Auth Token "Go-API"

func makeBugsnagDAAGet(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "token "+DAA_TOKEN)

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
