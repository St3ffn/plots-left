// cli tool to find out how many plots will still fit on your hard disk
package main

import (
	"fmt"
	"github.com/St3ffn/plots-left/internal/cli"
	"github.com/St3ffn/plots-left/internal/printer"
	"github.com/St3ffn/plots-left/pkg/disk"
	"os"
)

var (
	stderr = os.Stderr
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(stderr, "%s\n", err)
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

	var p printer.Printer = printer.DefaultPrinter{}
	if ctx.Verbose {
		p = printer.VerbosePrinter{}
	}

	p.Print(&info)

	return nil
}
