package github

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-github/v73/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSummarizeComparison(t *testing.T) {
	testCases := []struct {
		filename string
	}{
		{
			filename: "golang_go_example_comparison.json",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.filename, func(t *testing.T) {
			testFilePath := filepath.Join("testdata", tt.filename)
			data, err := os.ReadFile(testFilePath)
			require.NoError(t, err, "failed to read test data file")

			var comparison github.CommitsComparison
			err = json.Unmarshal(data, &comparison)
			require.NoError(t, err, "failed to unmarshal test data")

			summary := SummarizeComparison(&comparison)

			t.Logf("Summary: \n%s", summary)

			// ensure each field in the summary is not empty
			assert.NotEmpty(t, summary.URL, "URL")
			assert.Empty(t, summary.BehindBy, "BehindBy")
			assert.NotEmpty(t, summary.AheadBy, "AheadBy")
			assert.NotEmpty(t, summary.Commits, "Commits")
			for i, commit := range summary.Commits {
				assert.NotEmpty(t, commit.message, "commit message at index %d", i)
				assert.NotEmpty(t, commit.author, "commit author at index %d", i)
				assert.NotEmpty(t, commit.date, "commit date at index %d", i)
			}
		})
	}
}
