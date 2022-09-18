package selector

import (
	"errors"
	"fmt"
	"kcalc/internel/flag"
	"kcalc/pkg/conv"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v2"
)

var (
	errFormat = errors.New("format selector error")
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:   "selector",
		Usage:  "Parse Segment Selector",
		Action: action,
		Flags: []cli.Flag{
			&flag.ValueFlag,
		},
	}
}

func action(ctx *cli.Context) error {
	s, err := NewSelector(ctx.Int64(flag.ValueName))
	if err != nil {
		return err
	}

	s.FPrint()
	s.Description()
	return nil
}

type Selector struct {
	Origin int64
	Binary string

	RPL   int64
	TI    int64
	Index int64
}

func NewSelector(o int64) (*Selector, error) {
	bin := conv.Int2Bin(o, 16)

	rpl, err := strconv.ParseInt(bin[14:], 2, 0)
	if err != nil {
		return nil, errFormat
	}

	ti, err := strconv.ParseInt(bin[13:14], 2, 0)
	if err != nil {
		return nil, errFormat
	}

	index, err := strconv.ParseInt(bin[:13], 2, 0)
	if err != nil {
		return nil, errFormat
	}

	return &Selector{
		Origin: o,
		Binary: bin,

		RPL:   rpl,
		TI:    ti,
		Index: index,
	}, nil
}

func (s *Selector) FPrint() {
	intb := table.NewWriter()

	intb.SetTitle(fmt.Sprintf("# Input: %d, Binary: %s", s.Origin, s.Binary))
	intb.SetOutputMirror(os.Stdout)
	intb.Style().Options.SeparateRows = true
	intb.Style().Options.SeparateColumns = true
	intb.Style().Options.DrawBorder = true

	intb.AppendHeader(table.Row{
		"Index", "Member", "Value", "Mean",
	})

	intb.AppendRows([]table.Row{
		{"0~1", "RPL", s.RPL, "Requested Privilege Level"},
		{"2", "TI", s.TI, "Table Indicator (0:GDT;1:LDT)"},
		{"3~15", "Index", s.Index, "Segment Descriptor Address = GDT/IDT Base Address + 8*Index"},
	})

	intb.Render()
}

func (s *Selector) Description() {
	destb := table.NewWriter()

	destb.SetTitle("Selector Structure")
	destb.SetOutputMirror(os.Stdout)
	destb.Style().Options.SeparateRows = true
	destb.Style().Title.Align = text.AlignCenter

	destb.AppendRows([]table.Row{
		{
			"15,14,13,12,11,10,9,8,7,6,5,4,3",
			"2", "1, 0",
		},
		{
			"Index", "TI", "RPL",
		},
	})

	destb.Render()
}
