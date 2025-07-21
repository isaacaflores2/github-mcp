package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v73/github"
)

// Client GitHub client wrapper
type Client struct {
	// client is the underlying GitHub client
	client *github.Client
}

// NewClient creates a new GitHub client using API key authentication
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	return &Client{
		client: github.NewClient(nil).WithAuthToken(apiKey),
	}, nil
}

// CompareCommits compares two commits, refs (branches or tags) or commit SHAs.
func (c *Client) CompareCommits(ctx context.Context, owner, repo, base, head string) (*github.CommitsComparison, error) {
	comparison, _, err := c.client.Repositories.CompareCommits(ctx, owner, repo, base, head, nil)
	if err != nil {
		return nil, err
	}

	return comparison, nil
}
