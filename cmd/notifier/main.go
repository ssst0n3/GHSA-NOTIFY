package main

import (
	"encoding/json"
	"ghsa-notify/client"
	"ghsa-notify/ghsa"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
)

func Common() {
	var packages []string
	awesome_error.CheckFatal(json.Unmarshal([]byte(os.Getenv("packages")), &packages))
	feedFilePath := os.Getenv("feed")
	token := os.Getenv("GHTOKEN")

	var securityVulnerabilities []ghsa.SecurityVulnerability
	for _, pkg := range packages {
		query := ghsa.ListSecurityVulnerabilitiesByPackage(client.New(token), pkg, 100)
		securityVulnerabilities = append(securityVulnerabilities, query.SecurityVulnerabilityConnection.Nodes...)
	}
	feed, err := ghsa.GenerateNewFeed(securityVulnerabilities)
	awesome_error.CheckFatal(err)
	equal, err := ghsa.CompareFeed(*feed, feedFilePath)
	awesome_error.CheckFatal(err)
	if !equal {
		err = ghsa.WriteRss(feed, feedFilePath)
		awesome_error.CheckFatal(err)
	}
}

func main() {
	Common()
}
