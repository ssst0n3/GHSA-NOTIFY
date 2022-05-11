package ghsa

import (
	"ghsa-notify/client"
	"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func TestListSecurityVulnerabilitiesByPackage(t *testing.T) {
	c := client.New(os.Getenv("GHTOKEN_TEST"))
	q := ListSecurityVulnerabilitiesByPackage(c, "github.com/docker/docker", 10)
	spew.Dump(q)
}

func TestListSecurityVulnerabilitiesByRepository(t *testing.T) {
	c := client.New(os.Getenv("GHTOKEN_TEST"))
	q := ListSecurityVulnerabilitiesByRepository(c, "containerd", "containerd", 10)
	spew.Dump(q)
}
