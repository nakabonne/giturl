package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

type sshOptions struct {
	scpLike bool
	user    string
}

func NewSSHCommand(stdout io.Writer) *cobra.Command {
	s := &sshOptions{}
	cmd := &cobra.Command{
		Use:     "ssh",
		Short:   "Convert into sshOptions syntax",
		Example: "giturl ssh --scp-like --user=git https://github.com/org/repo.git",
	}
	cmd.Flags().BoolVarP(&s.scpLike, "scp-like", "s", s.scpLike, "emit scp-like syntax")
	cmd.Flags().StringVar(&s.user, "user", s.user, "override the user")

	r := &runner{
		stdout:      stdout,
		scheme:      converter.SchemeSSH,
		makeOptions: s.makeOptions,
	}
	cmd.RunE = r.run

	return cmd
}

func (s *sshOptions) makeOptions() *converter.Options {
	opts := &converter.Options{
		ScpLike: s.scpLike,
	}
	if s.user != "" {
		opts.User = &s.user
	}
	return opts
}
