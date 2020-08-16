package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewGitCommand(stdout io.Writer) *cobra.Command {
	r := &runner{
		stdout: stdout,
		scheme: converter.SchemeGit,
	}
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Convert into git syntax",
		RunE:  r.run,
	}

	return cmd
}
