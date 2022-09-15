package main

import (
	"kcalc/internel/app"
	"os"
)

func main() {
	app := app.NewApp()

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
