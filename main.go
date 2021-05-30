// cli tool to find out how many plots will still fit on your hard disk
package main

import (
	"fmt"
	"github.com/St3ffn/plots-left/internal/cli"
	"github.com/St3ffn/plots-left/pkg/disk"
	"os"
)

var (
	stdout = os.Stdout
	stderr = os.Stderr
)

func main() {
	left, err := run()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(stdout, left)
}

// run the cli
func run() (uint64, error) {
	ctx, err := cli.RunCli()
	if err != nil {
		return 0, err
	}
	if ctx.Done {
		return 0, nil
	}
	d, err := disk.NewDisk(ctx.Path)
	if err != nil {
		return 0, err
	}
	info := disk.PlotInfo{Disk: d, Reserved: ctx.Reserved}

	return info.PlotsLeft(), nil
}
