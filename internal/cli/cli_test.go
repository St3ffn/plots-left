package cli

import (
	"errors"
	"reflect"
	"testing"
)

func TestRunCli(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Context
		wantErr error
	}{
		{
			name: "ok",
			args: []string{"plots-left", "/my/fancy/path"},
			want: Context{
				Reserved: Reserved,
				Path:     "/my/fancy/path",
				Done:     false,
			},
		},
		{
			name: "help short",
			args: []string{"plots-left", "-h"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "help long",
			args: []string{"plots-left", "--help"},
			want: Context{
				Reserved: Reserved,
				Done:     true,
			},
		},
		{
			name: "reserve none short",
			args: []string{"plots-left", "-r", "0", "/my/fancy/path"},
			want: Context{
				Reserved: 0,
				Path:     "/my/fancy/path",
				Done:     false,
			},
		},
		{
			name: "reserve 11231230 long",
			args: []string{"plots-left", "--reserve", "11231230", "/my/fancy/path"},
			want: Context{
				Reserved: 11231230,
				Path:     "/my/fancy/path",
				Done:     false,
			},
		},
		{
			name:    "err no path",
			args:    []string{"plots-left"},
			wantErr: errors.New("PATH is missing"),
		},
		{
			name:    "err too many paths",
			args:    []string{"plots-left", "/tmp/a", "other/path"},
			wantErr: errors.New("only one PATH allowed"),
		},
		{
			name:    "unknown parameter -x",
			args:    []string{"plots-left", "-x", "asdas"},
			wantErr: errors.New("flag provided but not defined: -x"),
		},
		{
			name:    "invalid reserve paramter",
			args:    []string{"plots-left", "-r", "12.12", "/home/steffen"},
			wantErr: errors.New("invalid value \"12.12\" for flag -r: parse error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Args = tt.args
			got, err := RunCli()
			if err != nil {
				if tt.wantErr == nil || !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("RunCli() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("RunCli() got = %v, want %v", got, tt.want)
			}
		})
	}
}
