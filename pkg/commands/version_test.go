package commands

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersionCommand(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		commit     string
		date       string
		wantStdout string
		want       string
	}{
		{
			name: "nothing populated",
			want: fmt.Sprintf("version=unversioned, commit=?, buildDate=?, os=%s, arch=%s\n", runtime.GOOS, runtime.GOARCH),
		},
		{
			name:    "1.0.0",
			version: "1.0.0",
			commit:  "abc",
			date:    "2020-01-01T00:00:20Z",
			want:    fmt.Sprintf("version=1.0.0, commit=abc, buildDate=2020-01-01T00:00:20Z, os=%s, arch=%s\n", runtime.GOOS, runtime.GOARCH),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			if tt.version != "" {
				version = tt.version
			}
			if tt.commit != "" {
				commit = tt.commit
			}
			if tt.date != "" {
				date = tt.date
			}

			cmd := NewVersionCommand(b)
			cmd.Run(nil, []string{})
			assert.Equal(t, tt.want, b.String())
		})
	}
}
