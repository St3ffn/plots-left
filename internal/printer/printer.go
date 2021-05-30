// Package printer contains functionality to write output
package printer

import (
	"fmt"
	"github.com/St3ffn/plots-left/pkg/disk"
	"os"
	"text/tabwriter"
)

var Output = os.Stdout

// Printer interface to print something
type Printer interface {
	Print(info *disk.PlotInfo)
}

type DefaultPrinter struct{}

// Print prints the amount of plots left
func (d DefaultPrinter) Print(info *disk.PlotInfo) {
	_, _ = fmt.Fprintln(Output, info.PlotsLeft())
}

type VerbosePrinter struct {}

const verboseHeader string = "Path\tTotal\tStored\tReserved\tLeft\n"

// Print prints two lines with the following columns:
// Path, Total Amount of Plots, Amount of Plots Stored, Plots Left
func (v VerbosePrinter) Print(info *disk.PlotInfo) {
	w := tabwriter.NewWriter(Output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprint(w, verboseHeader)
	_, _ = fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\n",
		info.Path, info.PlotsTotal(), info.PlotsStored(), info.Reserved, info.PlotsLeft())
	_ = w.Flush()
}
