package commands

import (
	"fmt"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	)

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
