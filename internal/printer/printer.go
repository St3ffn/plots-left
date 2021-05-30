// Package printer contains functionality to write output
package printer

import (
	"fmt"
	"github.com/St3ffn/plots-left/pkg/disk"
	"os"
)

var Output = os.Stdout

// Printer interface to print something
type Printer interface {
	Print(info *disk.PlotInfo)
}

type DefaultPrinter struct {}

func (d DefaultPrinter) Print(info *disk.PlotInfo) {
	_, _ = fmt.Fprintln(Output, info.PlotsLeft())
}

type VerbosePrinter struct {}

func (v VerbosePrinter) Print(info *disk.PlotInfo) {
	_, _ = fmt.Fprintf(Output, "Path\tTotal\tReserved\tLeft\n")
	_, _ = fmt.Fprintf(Output, "%s\t%d\t%d\t\t%d\n", info.Path, info.PlotsTotal(), info.Reserved, info.PlotsLeft())
}
