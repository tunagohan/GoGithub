package main

import (
	"context"
	"log"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "{...Your Access Token...}"},
	)
	tc := oauth2.NewClient(ctx, ts)
	cl := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "{...UserName...}", nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(repos)
}
