package github

import "github.com/msolimans/gitcomp/pkg/build"

const (
	defaultBaseURL   = "https://api.github.com/"
	acceptHeader              = "application/vnd.github+json"  //chk https://docs.github.com/en/rest/overview/media-types
)
var defaultUserAgent = "gitcomp-" + build.Revision