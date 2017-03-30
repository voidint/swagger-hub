package cli

import "github.com/urfave/cli"

var (
	commands = []cli.Command{
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Run document server",
			Flags:     []cli.Flag{flDir, flDomain, flPort},
			Action:    run,
		},
	}
)
