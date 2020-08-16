package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nakabonne/giturl/pkg/converter"
)

func TestRunnerRun(t *testing.T) {
	tests := []struct {
		name    string
		scheme  converter.Scheme
		args    []string
		wantErr bool
	}{
		{
			name:    "no args given",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "unconvertable scheme",
			scheme:  converter.Scheme("foo"),
			args:    []string{"foo://path/to/repo.git"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		b := new(bytes.Buffer)
		r := &runner{
			stdout: b,
			scheme: tt.scheme,
		}
		err := r.run(nil, tt.args)
		assert.Equal(t, tt.wantErr, err != nil)
	}
}
