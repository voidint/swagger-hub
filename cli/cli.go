package cli

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli"
	"github.com/voidint/swagger-hub/build"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

// Run run cli
func Run() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "Open API documents server"
	app.Version = build.Version("1.0.1")
	app.Authors = []cli.Author{
		{Name: "voidint", Email: "voidint@126.com"},
	}

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
