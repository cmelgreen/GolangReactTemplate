package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)

type Path struct {
	Path 		string 	`json:"path"`
	Redirect 	bool 	`json:"redirect"`
}

type RouteFile struct {
	Routes		[]*Path	`json:"routes"`
}

func parseRoutes(file string) []string {
	jsonFile, err := os.Open(file)
	if err != nil {
		return []string{}
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile) 
	if err != nil {
		return []string{}
	}

	fmt.Println(string(jsonBytes))

	var routes RouteFile
	json.Unmarshal(jsonBytes, &routes)

	fmt.Println(routes)

	var pathsToRedirect []string

	for _, path := range routes.Routes {
		if path.Redirect == true {
			pathsToRedirect = append(pathsToRedirect, path.Path)
		}
	}

	fmt.Println(pathsToRedirect)

	return pathsToRedirect

}

