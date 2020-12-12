package commands

import (
	"strings"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	)

type DeleteCommand struct {
	Arg1, Arg2 string
}

func (d *DeleteCommand) Execute(loop engine.Handler) {
	result := strings.Replace(d.Arg1, d.Arg2, "", -1)
	loop.Post(&PrintCommand{Arg: result})
}