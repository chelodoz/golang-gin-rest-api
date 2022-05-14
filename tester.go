package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main1() {

	url := "http://localhost:8080/videos"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
