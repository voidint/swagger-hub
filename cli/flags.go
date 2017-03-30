package cli

import "github.com/urfave/cli"

var (
	// flLog = cli.StringFlag{
	// 	Name:   "log, l",
	// 	Usage:  "Log destination(file path or 'stdout')",
	// 	EnvVar: "SWAGGER_HUB_LOG",
	// }

	flPort = cli.UintFlag{
		Name:   "port, p",
		Usage:  "Document service port",
		Value:  8080,
		EnvVar: "SWAGGER_HUB_PORT",
	}

	flDomain = cli.StringFlag{
		Name:   "domain, D",
		Usage:  "Documents service domain",
		Value:  "localhost",
		EnvVar: "SWAGGER_HUB_DOMAIN",
	}

	flDir = cli.StringFlag{
		Name:   "dir, d",
		Usage:  "UI directory",
		EnvVar: "SWAGGER_HUB_DIR",
	}
)
