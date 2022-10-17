package github

import (
	"strconv"
	"strings"
)

func (c *CommitsComparison) String() string {
	var sb strings.Builder
	sb.WriteString("\n========================== Summary ==========================\n")
	sb.WriteString("\nUrl: " + c.URL + "\n HTMLUrl: " + c.HTMLURL + "\nDifferenceUrl: " + c.DiffURL + "\nPatchUrl: " + c.PatchURL)
	sb.WriteString("\nTotal commits: " + strconv.Itoa(c.TotalCommits))
	sb.WriteString("\nStatus: " + c.Status + " by " + map[string]string{"ahead": strconv.Itoa(c.AheadBy), "behind": strconv.Itoa(c.BehindBy)}[c.Status] + " commits \n")

	sb.WriteString("\n========================== Commits ==========================\n")
	for _,c := range c.Commits {
		sb.WriteString(c.String())
	}
	sb.WriteString("\n========================== File Changes ==========================\n")
	for _,f := range c.Files {
		sb.WriteString(f.String())
	}

	return sb.String()
}

func (c *CommitFile) String() string {
	return "Filename: " + c.Filename + "Status: " + c.Status + "Additions: " + strconv.Itoa(c.Additions) + "Deletions: " + strconv.Itoa(c.Deletions) + "Changes: " + strconv.Itoa(c.Changes) + "RawUrl: " + c.RawURL + "BlobUrl: " + c.BlobURL + "Patch: " + c.Patch
}

func (r *RepositoryCommit) String() string {
	return "\nSHA: " + r.Sha + "\nAuthor: " + r.Author.String() + "\nCommitter: " + r.Committer.String() + "\nMessage: " + r.Commit.Message + "\n"
}

func (u  *User) String() string {
	return u.Login
}

