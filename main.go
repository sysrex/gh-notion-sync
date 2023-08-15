package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Please set the GITHUB_TOKEN environment variable")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.IssueListOptions{
		Filter: "assigned",
		State:  "open",
	}

	issues, _, err := client.Issues.List(ctx, true, opt)
	if err != nil {
		log.Fatalf("Error listing assigned issues: %v", err)
	}

	fmt.Printf("Found %d issues assigned to you:\n", len(issues))

	for _, issue := range issues {
		fmt.Printf("#%d - %s\n", issue.GetNumber(), issue.GetTitle())
	}
}
