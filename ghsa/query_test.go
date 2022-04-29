package ghsa

import (
	"ghsa-notify/client"
	"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	c := client.New(os.Getenv("GHTOKEN_TEST"))
	q := ListSecurityVulnerabilities(c, "github.com/docker/docker", 10)
	spew.Dump(q)
}
