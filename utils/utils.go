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
	req.Header.Set("X-Riot-Token", "RGAPI-751a12c2-39f9-464e-a489-c6a58ee2c0f3")
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
