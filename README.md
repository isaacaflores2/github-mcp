# github-mcp
Simple mcp server for GitHub related things

## Tools
- `compare_commits` provides a summary of the changes between two commits/refs

## Setup
1. Install the server
```bash
  go install ./cmd/githubmcp 
```

2. Configure your MCP client to connect to your server
### Windsurf
Update `~/.codeium/windsurf/mcp_config.json` to include the following:
```json
{
  "mcpServers": {
    "github-mcp": {
      "command": "./full/path/go/bin/githubmcp",
      "args": [],
      "env": {
        "GITHUB_TOKEN": "<token>"
      }
    }
  }
}
```

## Usage
<details>
<summary>Compare Commits Example</summary>

Example prompt which will generate a changelog entry and saves the comparison summary to a file:
```
Can you compare commits base: 0de8565 to head: c4c8a40? owner: isaacaflores2, repo: github-mcp? 
Save the raw summary returned my the tool to a file with format <owner>_<repo>_<base>_<head> .
Then add a new entry to CHANGELOG.md. Add each commit to the changelog. I would like all commits to be added to the changelog even if they are automated version bumps. Each commit should be briefly summarized in the changelog
```

Generated changelog:
```markdown
# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased] - 2025-07-21

### Added
- Added commit_comparison tool which summarizes the output of github compare (8c997cc)
- Added tools section to README.md (56e3163)

### Changed
- Updated README.md (ca9d302)
- Fixed README.md formatting (95d61ce)
```

Saved output:
```
Comparison URL: https://api.github.com/repos/isaacaflores2/github-mcp/compare/0de8565...56e3163 
Base is behind by 0 commits and ahead by 4 
Here is a summary of the commits: 
 
-----------Commit-----------
Author: Isaac Flores
Date: 2025-07-21 15:13:02 -0700 PDT 
Message: added commit_comparison tool which summarizes the output of github compare
URL: https://github.com/isaacaflores2/github-mcp/commit/8c997cc29bc7a77450413a4500508f9894f7ff4d

-----------Commit-----------
Author: Isaac Flores
Date: 2025-07-21 15:26:50 -0700 PDT 
Message: updated README.md
URL: https://github.com/isaacaflores2/github-mcp/commit/ca9d3028c1a82776ce9c12a1ec43fc8cc21ad7f8

-----------Commit-----------
Author: Isaac Flores
Date: 2025-07-21 17:20:14 -0700 PDT 
Message: fixed README.md formatting
URL: https://github.com/isaacaflores2/github-mcp/commit/95d61cea71b593ba699e9dcc5347a3020cc930b8

-----------Commit-----------
Author: Isaac Flores
Date: 2025-07-21 17:21:02 -0700 PDT 
Message: added tools section to README.md
URL: https://github.com/isaacaflores2/github-mcp/commit/56e31630fa7e2c2113e2b973aa5ed03a72eddea8
```

</details>
