package commands

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

type ssh struct {
	scpLike bool
	user    string
	stdout  io.Writer
}

func NewSSHCommand(stdout io.Writer) *cobra.Command {
	r := &ssh{
		stdout: stdout,
	}
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "Convert into ssh syntax",
		RunE:  r.run,
	}
	cmd.Flags().BoolVarP(&r.scpLike, "scp-like", "s", r.scpLike, "emit scp-like syntax")
	cmd.Flags().StringVar(&r.user, "user", r.user, "override the user")

	return cmd
}

func (r *ssh) run(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no URL given")
	}
	opts := converter.Options{
		ScpLike: r.scpLike,
	}
	if r.user != "" {
		opts.User = &r.user
	}
	res, err := converter.Convert(args[0], converter.SchemeSSH, opts)
	if err != nil {
		return err
	}
	fmt.Fprintln(r.stdout, res)
	return nil
}
