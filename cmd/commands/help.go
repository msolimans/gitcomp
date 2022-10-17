package commands

import (
	"fmt"

	"github.com/msolimans/gitcomp/pkg/build"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
 
	command := &cobra.Command{
		Use:   "version",
		Short: "Version of the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Time:\t", build.Time)
			fmt.Println("Revision:\t", build.Revision)
		},
	}

	return command
}