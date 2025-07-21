package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/isaacaflores2/github-mcp/internal/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type toolHandlers struct {
	gitHubClient *github.Client
}

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"GitHub MCP Server",
		"0.1.0",
		server.WithToolCapabilities(false),
	)

	client, err := github.NewClient(os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	handlers := &toolHandlers{
		gitHubClient: client,
	}

	// Add tools
	compareCommits := mcp.NewTool("compare_commits",
		mcp.WithDescription("Compares two commits against one another and returns a summary. You can compare refs (branches or tags) and commit SHAs in the same repository"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("Owner of the repository"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("Name of the repository"),
		),
		mcp.WithString("base",
			mcp.Required(),
			mcp.Description("Base commit or base ref (e.g. a branch name or a commit SHA)"),
		),
		mcp.WithString("head",
			mcp.Required(),
			mcp.Description("Head commit or head ref (e.g. a branch name or a commit SHA)"),
		),
	)

	// Add tool handlers
	s.AddTool(compareCommits, handlers.compareCommitsHandler)

	// Start the stdio server
	fmt.Println("Starting github mcp server!")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func (th *toolHandlers) compareCommitsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner, err := request.RequireString("owner")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	repo, err := request.RequireString("repo")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	base, err := request.RequireString("base")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	head, err := request.RequireString("head")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	comparison, err := th.gitHubClient.CompareCommits(ctx, owner, repo, base, head)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(github.SummarizeComparison(comparison).String()), nil
}
