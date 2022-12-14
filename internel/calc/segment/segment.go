package segment

import (
	"kcalc/internel/flag"
	"kcalc/pkg/conv"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:   "segment",
		Usage:  "Parse Segment Descriptor",
		Action: action,
		Flags: []cli.Flag{
			&flag.AddrFlag,
		},
	}
}

func action(ctx *cli.Context) error {
	s, err := NewSegment(ctx)
	if err != nil {
		return err
	}

	s.FPrint()
	s.Description()

	return nil
}

type Segment struct {
	Origin string
	Binary *conv.Address
}

func NewSegment(ctx *cli.Context) (*Segment, error) {
	ori := ctx.String(flag.AddrName)

	bin, err := conv.Addr64ToBin(ori)
	if err != nil {
		return nil, err
	}

	return &Segment{
		Origin: ori,
		Binary: bin,
	}, nil
}

func (s *Segment) FPrint() {}

func (s *Segment) Description() {
	htb := table.NewWriter()
	htb.SetTitle("Segment Descriptor (High 32-bits)")
	htb.SetOutputMirror(os.Stdout)
	htb.Style().Options.SeparateRows = true
	htb.Style().Title.Align = text.AlignCenter
	htb.AppendRows([]table.Row{
		{
			"31,30,29,28,27,26,25,24",
			"23", "22", "21", "20",
			"19,18,17,16",
			"15", "14,13", "12",
			"11,10,9,8",
			"7,6,5,4,3,2,1,0",
		},
		{
			"Base 31:24",
			"G", "D/B", "0", "AVL",
			"Seg.Limit 19:16",
			"P", "DPL", "S",
			"TYPE",
			"Base 23:16",
		},
	})
	htb.Render()

	ltb := table.NewWriter()
	ltb.SetTitle("Segment Descriptor (Low 32-bits)")
	ltb.SetOutputMirror(os.Stdout)
	ltb.Style().Options.SeparateRows = true
	ltb.Style().Title.Align = text.AlignCenter
	ltb.AppendRows([]table.Row{
		{
			"31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16",
			"15,14,13,12,11,10,9,8,7,6,5,4,3,2,1",
		},
		{
			"Base 15:00",
			"Seg.Limit 15:00",
		},
	})
	ltb.Render()
}
