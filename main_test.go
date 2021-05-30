package main

import (
	"github.com/St3ffn/plots-left/internal/cli"
	"github.com/St3ffn/plots-left/internal/printer"
	"github.com/St3ffn/plots-left/pkg/disk"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		statfs  syscall.Statfs_t
		want    string
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
			want:    "20\n",
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
			want:    "0\n",
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
			want:    "1\n",
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
			tmpFile, err := ioutil.TempFile(os.TempDir(), "test-output")
			if err != nil {
				log.Fatal("Cannot create temporary file", err)
			}
			printer.Output = tmpFile

			gotErr := run()
			// Close the file
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}

			content, err := ioutil.ReadFile(tmpFile.Name())
			if err != nil {
				log.Fatal(err)
			}

			// Convert []byte to string and print to screen
			got := string(content)

			// cleanup
			_ = os.Remove(tmpFile.Name())

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
