package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/commands"
)

var (
	// Automatically populated by goreleaser during build
	version = "unversioned"
	commit  = "?"
	date    = "?"
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
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print the current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(a.stderr, "version=%s, commit=%s, buildDate=%s, os=%s, arch=%s\n", version, commit, date, runtime.GOOS, runtime.GOARCH)
		},
	}
	a.rootCmd.AddCommand(versionCmd)
	return a
}

func (a *app) addCommands(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		a.rootCmd.AddCommand(cmd)
	}
}

func main() {
	a := newApp("giturl", "A converter for Git URLs")

	a.addCommands(
		commands.NewSSHCommand(),
	)
	if err := a.run(); err != nil {
		fmt.Fprintln(a.stderr, err)
		os.Exit(1)
	}
}

func (a *app) run() error {
	return a.rootCmd.Execute()
}
