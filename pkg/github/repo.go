package github

import (
	"context"
	"fmt"
	"net/url"
)

//easier to add more services that interact with github (separation of concerns)
type repoService struct {
	client *GitClient 
}

func NewRepoService(client *GitClient) *repoService {
	return &repoService{client}
}

// https://docs.github.com/en/rest/commits/commits#compare-two-commits
func (s *repoService) CompareCommits(ctx context.Context, owner, repo string, base, head string) (*CommitsComparison, error) {
	escapedBase := url.QueryEscape(base)
	escapedHead := url.QueryEscape(head)

	u := fmt.Sprintf("repos/%v/%v/compare/%v...%v", owner, repo, escapedBase, escapedHead)

	req, err := s.client.InitRequest("GET", u)
	if err != nil {
		return nil, err
	}

	comp := &CommitsComparison{}
	err = s.client.Send(ctx, req, comp)
	if err != nil {
		return nil, err
	}

	return comp, nil
}
