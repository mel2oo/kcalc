package flag

import "github.com/urfave/cli/v2"

var (
	ValueName = "value"
	ValueFlag = cli.Int64Flag{
		Name:  ValueName,
		Value: 0,
		Usage: "input a value",
	}

	AddrName = "addr"
	AddrFlag = cli.StringFlag{
		Name:  AddrName,
		Value: "00000000`00000000",
		Usage: "input a address (64-bits)",
	}

	AllName = "all"
	AllFlag = cli.BoolFlag{
		Name:  AllName,
		Value: true,
		Usage: "show structure descriptor",
	}
)
