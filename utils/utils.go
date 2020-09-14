package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetReq does the obvious
func GetReq(endpoint string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("X-Riot-Token", "RGAPI-6cf6cccc-02be-48f4-9c2c-21a18ac86565")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Could not execute request", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Could not read body from response", err)
	}
	return body
}
