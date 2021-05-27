package cli

import (
	"errors"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

var (
	args            = os.Args
	reserved uint64 = 1
)

// Context describes the environment of the tool execution
type Context struct {
	// Reserved represents the amount of plots to be reserved
	Reserved uint64
	// the path
	Path string
	// Done indicates that we are done (--help, --version...)
	Done bool
}

// RunCli starts the cli which includes validation of parameters.
func RunCli() (*Context, error) {
	var path string
	var done bool

	cli.HelpFlag = &cli.BoolFlag{
		Name:        "help",
		Aliases:     []string{"h"},
		Usage:       "show help",
		Destination: &done,
	}

	app := &cli.App{
		Name:                 "plots-left",
		Usage:                "find out how many plots will still fit on your hard disk",
		UsageText:            "plots-left [-r RESERVE] PATH\n\t plots-left -r 1 /plots/nas1",
		ArgsUsage:            "PATH",
		Description:          "Tool will perform the following calculation (AVAILABLE_DISK_SPACE/SINGLE_PLOT_SIZE)-RESERVED_PLOTS.",
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:        "reserve",
				Aliases:     []string{"r"},
				Required:    false,
				Value:       reserved,
				DefaultText: strconv.FormatUint(reserved, 10),
				Usage:       "`RESERVE`. the amount of plots to reserve.",
				Destination: &reserved,
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

	err := app.Run(args)
	if err != nil {
		return nil, err
	}

	return &Context{
		Reserved: reserved,
		Path:     path,
		Done:     done,
	}, nil
}
