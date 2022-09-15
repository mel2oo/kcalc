package selector

import (
	"errors"
	"fmt"
	"kcalc/internel/flag"
	"strconv"

	"github.com/urfave/cli/v2"
)

var (
	errFormat = errors.New("format selector error")
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:   "selector",
		Usage:  "Parse Selector",
		Action: action,
		Flags: []cli.Flag{
			&flag.ValueFlag,
		},
	}
}

func action(ctx *cli.Context) error {
	va := ctx.Int64(flag.ValueName)
	bva := align(strconv.FormatInt(va, 2))

	ser, err := parse(bva)
	if err != nil {
		return err
	}

	return ser.FPrint()
}

const (
	sLen = 16
)

func align(str string) (s string) {
	l1 := sLen - len(str)
	for l1 > 0 {
		s += "0"
		l1--
	}

	return s + str
}

type Selector struct {
	RPL   string
	TI    string
	Index string
}

func parse(s string) (*Selector, error) {
	if len(s) != sLen {
		return nil, fmt.Errorf("selector length error: %d", len(s))
	}

	return &Selector{
		RPL:   s[14:],
		TI:    s[13:14],
		Index: s[:13],
	}, nil
}

func (s *Selector) FPrint() error {
	rpl, err := strconv.ParseInt(s.RPL, 2, 0)
	if err != nil {
		return errFormat
	}

	ti, err := strconv.ParseInt(s.TI, 2, 0)
	if err != nil {
		return errFormat
	}

	index, err := strconv.ParseInt(s.Index, 2, 0)
	if err != nil {
		return errFormat
	}

	fmt.Printf("RPL: %d\nTI: %d\nIndex: %d\n", rpl, ti, index)

	return nil
}
