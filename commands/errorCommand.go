package commands

import (
	"fmt"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
)

type ErrorCommand struct {
	ErrorStr string
}

func (e *ErrorCommand) Execute(loop engine.Handler) {
	result := fmt.Sprintf("SYNTAX ERROR: %s", e.ErrorStr)
	loop.Post(&PrintCommand{Arg: result})
}