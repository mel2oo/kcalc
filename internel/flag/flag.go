package flag

import "github.com/urfave/cli/v2"

var (
	ValueName = "value"
	ValueFlag = cli.Int64Flag{
		Name:  ValueName,
		Value: 0,
		Usage: "input a value",
	}

	AllName = "all"
	AllFlag = cli.BoolFlag{
		Name:  AllName,
		Value: true,
		Usage: "show structure descriptor",
	}
)
