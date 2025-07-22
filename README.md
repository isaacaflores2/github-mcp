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