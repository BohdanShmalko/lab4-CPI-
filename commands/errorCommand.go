package commands

import (
	"fmt"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
)

type errorCommand struct {
	errorStr string
}

func (e *errorCommand) Execute(loop engine.Handler) {
	result := fmt.Sprintf("SYNTAX ERROR: %s", e.errorStr)
	loop.Post(&printCommand{arg: result})
}