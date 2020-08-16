package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFileCommand(t *testing.T) {
	tests := []struct {
		name    string
		argURL  string
		wantURL string
		wantErr bool
	}{
		{
			name:    "Converting of ssh",
			argURL:  "ssh://github.com/org/repo.git",
			wantURL: "file:///org/repo.git\n",
			wantErr: false,
		},
		{
			name:    "Converting of ssh with user",
			argURL:  "ssh://user@github.com/org/repo.git",
			wantURL: "file:///org/repo.git\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		b := new(bytes.Buffer)
		cmd := NewFileCommand(b)

		err := cmd.RunE(nil, []string{tt.argURL})
		assert.Equal(t, tt.wantErr, err != nil)
		assert.Equal(t, tt.wantURL, b.String())
	}
}
