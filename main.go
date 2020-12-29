package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// This will store all of the dioDB data in memory, updated reguarly
var cache string

func updateCache() {
	url := "https://raw.githubusercontent.com/disclose/diodb/master/program-list/program-list.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Retrieving data failed: ", err)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Update the cache
	cache = string(body)
}

func serveData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, cache)
}

func main() {

	// Update the cache every 60 minutes
	go func() {
		for {
			updateCache()
			time.Sleep(60 * time.Minute)
		}
	}()

	// Serve it as JSON
	http.HandleFunc("/", serveData)
	http.ListenAndServe(":80", nil)

}
