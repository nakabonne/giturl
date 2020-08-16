package commands

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSSHCommand(t *testing.T) {
	tests := []struct {
		name    string
		argURL  string
		scpLike bool
		user    string
		wantURL string
		wantErr bool
	}{
		{
			name:    "Converting of https",
			argURL:  "https://github.com/org/repo.git",
			wantURL: "ssh://github.com/org/repo.git\n",
			wantErr: false,
		},
		{
			name:    "Converting ssh into SCP-like syntax",
			argURL:  "https://github.com/org/repo.git",
			scpLike: true,
			wantURL: "github.com:org/repo.git\n",
			wantErr: false,
		},
		{
			name:    "Converting ssh into SCP-like syntax while overriding user",
			argURL:  "https://github.com/org/repo.git",
			scpLike: true,
			user:    "git",
			wantURL: "git@github.com:org/repo.git\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		b := new(bytes.Buffer)
		cmd := NewSSHCommand(b)
		cmd.Flags().Set("scp-like", fmt.Sprintf("%v", tt.scpLike))
		cmd.Flags().Set("user", tt.user)

		err := cmd.RunE(nil, []string{tt.argURL})
		assert.Equal(t, tt.wantErr, err != nil)
		assert.Equal(t, tt.wantURL, b.String())
	}
}
