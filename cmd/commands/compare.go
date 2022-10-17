package commands

import (
	"github.com/msolimans/gitcomp/pkg/compare"
	"github.com/spf13/cobra"
)

func NewCompareCommand() *cobra.Command {
	var (
		token 		string	
		org 		string	
		repo 		string
		baseSha 	string
		headSha    	string
	)
	
	command := &cobra.Command{
		Use:   "compare",
		Short: "Compare two commits",
		Run: func(cmd *cobra.Command, args []string) {
			compare.Start(token, org, repo, baseSha, headSha)
		},
	}

	//define flags
	command.Flags().StringVarP(&token, "token", "t", "", "personal github token")
	command.Flags().StringVarP(&org, "org", "o", "", "organization name or owner")
	command.Flags().StringVarP(&repo, "repo", "r", "", "repository name")
	command.Flags().StringVarP(&baseSha, "base", "b", "", "base commit sha")
	command.Flags().StringVarP(&headSha, "head", "d", "", "head commit sha")

	//mark flags as required
	command.MarkFlagRequired("token")
	command.MarkFlagRequired("org")
	command.MarkFlagRequired("repo")
	command.MarkFlagRequired("head")
	command.MarkFlagRequired("base")
	return command
}
