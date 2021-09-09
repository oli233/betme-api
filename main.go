package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiKey = "9317664c724de4ae8eb40bbb38a21d18"

type Data struct {
	Key    string `json:"key"`
	Active bool   `json:"active"`
	Group  string `json:"group"`
	Details string `json:"details"`
	Title string `json:"title"`
	Has_outrights bool `json:"has_outrights"`
}

type Sports struct {
	Success bool `json:"success"`
	DataArray []Data `json:"data"`
}

func main() {
	resp, err := http.Get("https://api.the-odds-api.com/v3/sports?apiKey="+ ApiKey)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	//decoding to json
	var jsonResult Sports

	//err = json.Unmarshal([]byte(string(body)), &jsonResult)
	err = json.Unmarshal([]byte(string(body)), &jsonResult)
	if err != nil {
		return
	}

	fmt.Println(jsonResult.DataArray)

	//if jsonResult["success"]==nil {
	//	log.Print("Didn't receive any data")
	//	return
	//} else {
	//	//print(jsonResult["data"])
	//}

	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)
}
