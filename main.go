// cli tool to find out how many plots will still fit on your hard disk
package main

import (
	"fmt"
	"github.com/St3ffn/plots-left/internal/cli"
	"github.com/St3ffn/plots-left/pkg/disk"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run the cli
func run() error {
	ctx, err := cli.RunCli()
	if err != nil {
		return err
	}
	if ctx.Done {
		return nil
	}
	d, err := disk.NewDisk(ctx.Path)
	if err != nil {
		return err
	}
	info := disk.PlotInfo{Disk: d, Reserved: ctx.Reserved}

	fmt.Println(info.PlotsLeft())

	return nil
}
