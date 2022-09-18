package app

import (
	"kcalc/internel/calc/segment"
	"kcalc/internel/calc/selector"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "kcalc"
	app.Version = "v1.0.0"
	app.Usage = "windows kernel calculator"
	app.Commands = []*cli.Command{
		selector.NewCommand(),
		segment.NewCommand(),
	}
	return app
}
