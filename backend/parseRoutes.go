package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Route struct {
	Path 		string 	`json:"path"`
	Redirect 	bool 	`json:"redirect"`
}

type RouteFile struct {
	Routes		[]*Route	`json:"routes"`
}

func parseRoutes(file string) []string {
	var routeFile RouteFile
	var routeStrings []string

	json.Unmarshal(readJSONFile(file), &routeFile)

	for _, route := range routeFile.Routes {
		if route.Redirect == true {
			routeStrings = append(routeStrings, route.Path)
		}
	}

	return routeStrings
}

func readJSONFile(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		return []byte{}
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile) 
	if err != nil {
		return []byte{}
	}

	return jsonBytes
}
