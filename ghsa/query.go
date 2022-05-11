package ghsa

import (
	"context"
	"github.com/shurcooL/githubv4"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func ListSecurityVulnerabilitiesByPackage(client *githubv4.Client, pkg string, first int) (query SecurityVulnerabilities) {
	err := client.Query(context.Background(), &query, map[string]interface{}{
		"package": githubv4.String(pkg),
		"first":   githubv4.Int(first),
	})
	awesome_error.CheckFatal(err)
	return
}

func ListSecurityVulnerabilitiesByRepository(client *githubv4.Client, name string, owner string, first int) (query QueryRepository) {
	err := client.Query(context.Background(), &query, map[string]interface{}{
		"name":  githubv4.String(name),
		"owner": githubv4.String(owner),
		"first": githubv4.Int(first),
	})
	awesome_error.CheckFatal(err)
	return
}
