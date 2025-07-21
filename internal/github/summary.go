package github

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v73/github"
)

// CompareSummary summary of a comparison
type CompareSummary struct {
	URL      string
	BehindBy int
	AheadBy  int
	Commits  ShortCommits
}

func (c CompareSummary) String() string {
	return fmt.Sprintf(
		"Comparison URL: %s \nBase is behind by %d commits and ahead by %d \nHere is a summary of the commits: \n %s \n",
		c.URL,
		c.BehindBy,
		c.AheadBy,
		c.Commits,
	)
}

type ShortCommits []ShortCommit

func (c ShortCommits) String() string {
	var s strings.Builder
	for _, commit := range c {
		s.WriteString(commit.String())
	}
	return s.String()
}

// ShortCommit brief summary of a commit
type ShortCommit struct {
	message string
	author  string
	date    string
	url     string
}

func (c ShortCommit) String() string {
	return fmt.Sprintf(
		"\n-----------Commit-----------\nAuthor: %s\nDate: %s \nMessage: %s\nURL: %s\n",
		c.author,
		c.date,
		c.message,
		c.url,
	)
}

func SummarizeComparison(comparison *github.CommitsComparison) CompareSummary {
	commits := make([]ShortCommit, len(comparison.Commits))
	for i, c := range comparison.Commits {
		c := ShortCommit{
			message: c.Commit.GetMessage(),
			author:  c.Commit.Author.GetName(),
			date:    c.Commit.Author.GetDate().Local().String(),
			url:     c.GetHTMLURL(),
		}
		commits[i] = c
	}
	return CompareSummary{
		URL:      comparison.GetURL(),
		BehindBy: comparison.GetBehindBy(),
		AheadBy:  comparison.GetAheadBy(),
		Commits:  commits,
	}
}
