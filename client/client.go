package client

import (
	"context"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func New(token string) *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}
