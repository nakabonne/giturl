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
		cmd     *cobra.Command
		want    string
		wantErr bool
	}{
		{
			name: "command to do nothing",
			cmd: &cobra.Command{
				Use: "sub",
				RunE: func(_ *cobra.Command, args []string) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			a := newApp("root", "desc", b, b)
			a.addCommands()

			err := a.run()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, b.String())
		})
	}
}
