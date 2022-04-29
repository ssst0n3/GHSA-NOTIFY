package main

import (
	"encoding/json"
	"ghsa-notify/client"
	"ghsa-notify/ghsa"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
)

func main() {
	var packages []string
	awesome_error.CheckFatal(json.Unmarshal([]byte(os.Getenv("packages")), &packages))
	for _, pkg := range packages {
		ghsa.ListSecurityVulnerabilities(client.New(os.Getenv("GHTOKEN")), pkg, 10)
	}
}
