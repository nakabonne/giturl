package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/commands"
)

type app struct {
	rootCmd *cobra.Command
	stdout  io.Writer
	stderr  io.Writer
}

func newApp(name, desc string) *app {
	a := &app{
		rootCmd: &cobra.Command{
			Use:   name,
			Short: desc,
		},
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	return a
}

func (a *app) addCommands(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		a.rootCmd.AddCommand(cmd)
	}
}

func (a *app) run() error {
	return a.rootCmd.Execute()
}

func main() {
	a := newApp("giturl", "Converts Git URLs into the scheme you like.")
	a.addCommands(
		commands.NewVersionCommand(a.stdout),
		commands.NewSSHCommand(a.stdout),
		commands.NewHTTPSCommand(a.stdout),
		commands.NewHTTPCommand(a.stdout),
		commands.NewGitCommand(a.stdout),
		commands.NewFileCommand(a.stdout),
	)

	if err := a.run(); err != nil {
		fmt.Fprintln(a.stderr, err)
		os.Exit(1)
	}
}
