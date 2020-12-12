package commands

import (
	"strings"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	)

type deleteCommand struct {
	arg1, arg2 string
}

func (d *deleteCommand) Execute(loop engine.Handler) {
	result := strings.Replace(d.arg1, d.arg2, "", -1)
	loop.Post(&printCommand{arg: result})
}