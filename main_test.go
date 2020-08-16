package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		subcmd  *cobra.Command
		want    string
		wantErr bool
	}{
		{
			name: "command to do nothing",
			subcmd: &cobra.Command{
				Use: "sub",
			},
			wantErr: false,
			want:    "desc\n\nUsage:\n\nFlags:\n  -h, --help   help for root\n\nAdditional help topics:\n  root sub    \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			a := newApp("root", "desc", b, b)
			a.rootCmd.SetOut(b)
			a.rootCmd.SetErr(b)
			a.addCommands(tt.subcmd)

			err := a.run()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, b.String())
		})
	}
}
