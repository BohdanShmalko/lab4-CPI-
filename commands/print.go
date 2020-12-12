package commands

import (
	"fmt"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	)

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}
