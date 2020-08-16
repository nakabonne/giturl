package commands

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunHTTPSCommand(t *testing.T) {
	tests := []struct {
		name    string
		argURL  string
		noUser  bool
		wantURL string
		wantErr bool
	}{
		{
			name:    "Converting of ssh",
			argURL:  "ssh://github.com/org/repo.git",
			noUser:  false,
			wantURL: "https://github.com/org/repo.git\n",
			wantErr: false,
		},
		{
			name:    "Converting of ssh while pruning user",
			argURL:  "ssh://user@github.com/org/repo.git",
			noUser:  true,
			wantURL: "https://github.com/org/repo.git\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		b := new(bytes.Buffer)
		cmd := NewHTTPSCommand(b)
		cmd.Flags().Set("no-user", fmt.Sprintf("%v", tt.noUser))

		err := cmd.RunE(nil, []string{tt.argURL})
		assert.Equal(t, tt.wantErr, err != nil)
		assert.Equal(t, tt.wantURL, b.String())
	}
}
