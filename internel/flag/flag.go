package flag

import "github.com/urfave/cli/v2"

var (
	ValueName = "value"
	ValueFlag = cli.Int64Flag{
		Name:  ValueName,
		Value: 0,
		Usage: "input a value",
	}
)
