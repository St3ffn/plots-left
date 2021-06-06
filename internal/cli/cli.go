// Package cli contains the command line interface, defines related parameters and validates them
package cli

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strconv"
)

var (
	Args = os.Args
	// Reserved defines the default amount of plots to reserve
	Reserved uint64 = 0
)

// Context describes the environment of the tool execution
type Context struct {
	// Reserved represents the amount of plots to be reserved
	Reserved uint64
	Path     string
	// Verbose mode
	Verbose bool
	// Done indicates that we are done (--help, --version...)
	Done bool
}

// RunCli starts the cli which includes validation of parameters.
func RunCli(writer io.Writer, version string) (*Context, error) {
	var path string
	var verbose, done bool

	cli.HelpFlag = &cli.BoolFlag{
		Name:        "help",
		Aliases:     []string{"h"},
		Usage:       "show help",
		Destination: &done,
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print version",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		_, _ = fmt.Fprintf(c.App.Writer, "%s version %s\n", c.App.Name, c.App.Version)
		done = true
	}

	app := &cli.App{
		Name:                 "plots-left",
		Usage:                "find out how many plots will still fit on your hard disk",
		UsageText:            "plots-left [-r RESERVE] [-v] PATH\n\t plots-left -v -r 1 /plots/nas1",
		ArgsUsage:            "PATH",
		Description:          "Tool will perform the following calculation (AVAILABLE_DISK_SPACE/SINGLE_PLOT_SIZE)-RESERVED_PLOTS.",
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Version:              version,
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:        "reserve",
				Aliases:     []string{"r"},
				Required:    false,
				Value:       Reserved,
				DefaultText: strconv.FormatUint(Reserved, 10),
				Usage:       "`RESERVE`. the amount of plots to reserve.",
				Destination: &Reserved,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Required:    false,
				Value:       false,
				Usage:       "enable verbose mode.",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return errors.New("PATH is missing")
			}
			if c.NArg() > 1 {
				return errors.New("only one PATH allowed")
			}

			path = c.Args().First()
			return nil
		},
		Copyright: "GNU GPLv3",
	}
	app.Writer = writer

	err := app.Run(Args)
	if err != nil {
		return nil, err
	}

	return &Context{
		Reserved: Reserved,
		Verbose:  verbose,
		Path:     path,
		Done:     done,
	}, nil
}
