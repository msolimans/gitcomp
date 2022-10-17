package compare

import (
	"context"
	"log"

	"github.com/msolimans/gitcomp/pkg/github"
	"github.com/msolimans/gitcomp/pkg/xhttp"
)

func Start(token, org, repo, baseSha, headSha string) {
	//	Sample request:
	// 	curl \
	//   -H "Accept: application/vnd.github+json" \
	//   -H "Authorization: Bearer <YOUR-TOKEN>" \
	//   https://github.com/github/linguist/compare/c3a414e..faf7c6f
	ctx := context.Background()
	
	//create http client with token
	httpClient := xhttp.NewHttpClient(ctx, token)
	//pass in http client to github client
	client := github.NewClient(httpClient)
	//create repo service
	repoSvc := github.NewRepoService(client)

	// https://api.github.com/repos/msolimans/Algorithms/compare/ff1912557b3ccb2e8f2c7d66f6317e7174763e2f...4d7d685f15d7187512c05ff4f42574a10f24554f
	// https://github.com/github/linguist/compare/c3a414e..faf7c6f
	comp,err := repoSvc.CompareCommits(ctx, org, repo, baseSha, headSha)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(comp.String())
}