package main

import (
	"bufio"
	"github.com/BohdanShmalko/lab4-CPI-/commands"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	"os"
	"strings"
)

func parse(strCommand string) engine.Command {
	parts := strings.Fields(strCommand)

	if len(parts) == 0 {
		return &commands.ErrorCommand{"empty command string"}
	}

	if parts[0] == "print" {
		if len(parts) < 2 {
			return &commands.ErrorCommand{"few arguments were introduced (must be 2 arguments)"}
		}
		return &commands.PrintCommand{parts[1]}
	} else if parts[0] == "delete" {
		if len(parts) < 3 {
			return &commands.ErrorCommand{"few arguments were introduced (must be 3 arguments)"}
		}
		return &commands.DeleteCommand{parts[1], parts[2]}
	}
	return &commands.ErrorCommand{"incorrect command name"}
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open("inputFile.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
