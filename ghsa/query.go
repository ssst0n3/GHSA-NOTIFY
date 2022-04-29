package ghsa

import (
	"context"
	"github.com/shurcooL/githubv4"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func ListSecurityVulnerabilities(client *githubv4.Client, pkg string, first int) (query SecurityVulnerabilities) {
	err := client.Query(context.Background(), &query, map[string]interface{}{
		"package": githubv4.String(pkg),
		"first":   githubv4.Int(first),
	})
	awesome_error.CheckFatal(err)
	return
}
