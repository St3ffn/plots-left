package main

import (
	"github.com/St3ffn/plots-left/internal/cli"
	"github.com/St3ffn/plots-left/pkg/disk"
	"syscall"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		statfs  syscall.Statfs_t
		want    uint64
		wantErr bool
	}{
		{
			name: "all free",
			args: []string{"plots-left", "/home/steffen"},
			statfs: syscall.Statfs_t{
				Bsize:  1,
				Blocks: uint64(disk.SizeOfPlot) * (20 + cli.Reserved),
				Bfree:  uint64(disk.SizeOfPlot) * (20 + cli.Reserved),
			},
			want:    20,
			wantErr: false,
		},
		{
			name: "none free",
			args: []string{"plots-left", "/home/steffen"},
			statfs: syscall.Statfs_t{
				Bsize:  1,
				Blocks: 12324234,
				Bfree:  12324230,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "one free",
			args: []string{"plots-left", "/home/steffen"},
			statfs: syscall.Statfs_t{
				Bsize:  1,
				Blocks: uint64(disk.SizeOfPlot) * (20 + cli.Reserved),
				Bfree:  uint64(disk.SizeOfPlot) * (1 + cli.Reserved),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk.Statfs = func(path string, stat *syscall.Statfs_t) (err error) {
				stat.Bsize = tt.statfs.Bsize
				stat.Blocks = tt.statfs.Blocks
				stat.Bfree = tt.statfs.Bfree
				return nil
			}
			cli.Args = tt.args
			got, err := run()
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
