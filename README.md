# github-mcp
Simple mcp server for GitHub related things

## Tools
- `compare_commits` provides a summary of the changes between two commits/refs

## Setup
1. Store your GitHub access token in the `GITHUB_TOKEN` environment variable
```bash
  export GITHUB_TOKEN=<your-token-here>
```
2. Install and run the server
```bash
  go install ./cmd/githubmcp 
  $GOBIN/githubmcp
```

3. Configure your MCP client to connect to your server
### Windsurf
Update `~/.codeium/windsurf/mcp_config.json` to include the following:
```json
{
  "mcpServers": {
    "github-mcp": {
      "command": "./full/path/go/bin/githubmcp",
      "args": [],
      "env": {}
    }
  }
}
```

## Usage
### Compare Commits
Example prompt which will generate a changelog entry and saves the comparison summary to a file:
```
Can you compare commits base: 0de8565 to head: c4c8a40? owner: isaacaflores2, repo: github-mcp? 
Save the raw summary returned my the tool to a file with format <owner>_<repo>_<base>_<head> .
Then add a new entry to CHANGELOG.md. Add each commit to the changelog. I would like all commits to be added to the changelog even if they are automated version bumps. Each commit should be briefly summarized in the changelog
```