package github

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)


const (
	//mock base url 
	baseURLPath = "/api"
)
	
func TestRepositoriesService_CompareCommits(t *testing.T) {
	testCases := []*testCase {
		prepareTestCase("b", "h"),
		prepareTestCase("`~!@#$%^&*()_+-90876ghj+12",  "123`~!@@#$%^&##$$}|;*-+"),
		prepare503TestCase(),
		prepareInvalidUrlTestCase(),
	}

	for _, sample := range testCases {
		gitCli, mux, shutdown := setup()
		 
		mux.HandleFunc(sample.pattern,sample.handler)

		ctx := context.Background()
		got, err := gitCli.CompareCommits(ctx, sample.org, sample.repo, sample.base, sample.head)
		if !sample.hasError && err != nil {
			t.Errorf("CompareCommits returned error: %v", err)
		}

		compareResponse(t, got, sample.want)

		shutdown()
	}
}

// test server for mocking along with a Client without passing any token to talk to that test server
func setup() (repoSvc *repoService, mux *http.ServeMux, teardown func()) {
	
	mux = http.NewServeMux()
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))

	server := httptest.NewServer(apiHandler)

	gitCli := NewClient(nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	// set the base url to the test server
	gitCli.BaseURL = url
	repoSvc = &repoService{client: gitCli}

	return repoSvc, mux, server.Close
}

type testCase struct {
	org string 
	repo string
	base string
	head string
	pattern string
	handler func(w http.ResponseWriter, r *http.Request)
	hasError 		bool
	want *CommitsComparison
}

func prepareTestCase(base, head string) *testCase {
	escapedBase := url.QueryEscape(base)
	escapedHead := url.QueryEscape(head)
	want := &CommitsComparison{
		Status:       "s",
		AheadBy:      1,
		BehindBy:     2,
		TotalCommits: 1,
		Commits: []*RepositoryCommit{
			{
					Sha: "s",
					Commit: &Commit{
						Author: &CommitAuthor{Name:  "n"},
					},
					Author:    &User{Login: "l"},
					Committer: &User{Login: "l"},
				},
			},
			Files: []*CommitFile{
				{
					Filename: "f",
				},
			},
			DiffURL:      fmt.Sprintf("https://github.com/o/r/compare/%v...%v.diff", escapedBase, escapedHead),
			URL:          fmt.Sprintf("https://api.github.com/repos/o/r/compare/%v...%v", escapedBase, escapedHead),
		}
	return &testCase {
			org: "o",
			repo: "r",
			base: base,
			head: head,
			want: want,
			pattern: fmt.Sprintf("/repos/%s/%s/compare/%s...%s", "o", "r", base, head),
			handler:  func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{
					"status": "s",
					"ahead_by": 1,
					"behind_by": 2,
					"total_commits": 1,
					"commits": [
						{
						"sha": "s",
						"commit": { "author": { "name": "n" } },
						"author": { "login": "l" },
						"committer": { "login": "l" }
						}
					],
					"files": [ { "filename": "f" } ],
					"diff_url":      "https://github.com/o/r/compare/%[1]v...%[2]v.diff",
					"url":           "https://api.github.com/repos/o/r/compare/%[1]v...%[2]v"}`, escapedBase, escapedHead)
			},
			}
}
func prepare503TestCase() *testCase {
	return &testCase {
		base: "b",
		head: "h",
		want: nil,
		hasError: true,
		pattern: fmt.Sprintf("/repos/%s/%s/compare/%s...%s", "o", "r", "b", "h"),
		handler:  func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
		},

	}
}


func compareResponse(t *testing.T, got *CommitsComparison, want *CommitsComparison) {
	t.Helper()
	if !cmp.Equal(got, want) {
		t.Errorf("Repositories.CompareCommits returned \n%+v, want \n%+v", got, want)
	}
}

func prepareInvalidUrlTestCase() *testCase {
	return &testCase {
		org: "\n",
		repo: "\n",
		head: "\n",
		base: "\n",
		hasError: true, 
		pattern: fmt.Sprintf("/repos/%s/%s/compare/%s...%s", "o", "r", "\n", "\n"),
		handler:  func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	}
}