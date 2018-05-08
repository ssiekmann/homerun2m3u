package main

// Written by Sascha Siekmann 2018
// Released into the public domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type GuideStruct []struct {
	GuideNumber string `json:"GuideNumber"`
	GuideName   string `json:"GuideName"`
	VideoCodec  string `json:"VideoCodec"`
	AudioCodec  string `json:"AudioCodec"`
	HD          int    `json:"HD,omitempty"`
	URL         string `json:"URL"`
}

func main() {
	argsWithoutProg := os.Args[1:]

	url := "http://" + argsWithoutProg[0] + "/lineup.json"

	Client := http.Client{Timeout: time.Second * 2} // Maximum of 2 secs

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := Client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	Channels := GuideStruct{}
	jsonErr := json.Unmarshal(body, &Channels)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("#EXTM3U")

	for i := 0; i < len(Channels); i++ {
		fmt.Println("#EXTINF:" + "-1," + Channels[i].GuideNumber + "," + Channels[i].GuideName)
		fmt.Println(Channels[i].URL)
	}
}
