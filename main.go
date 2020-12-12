package main

import (
	"bufio"
	"os"
	"strings"
	"github.com/BohdanShmalko/lab4-CPI-/engine"
	"github.com/BohdanShmalko/lab4-CPI-/commands"
)

func parse(strCommand string) engine.Command {
	parts := strings.Fields(strCommand)

	if len(parts) == 0 {
		res := new(commands.errorCommand)
		return res{"empty command string"}
	}

	if parts[0] == "print" {
		if len(parts) < 2 {
			res := new(commands.errorCommand)
			return res{"few arguments were introduced (must be 2 arguments)"}
		}
		res := new(commands.printCommand)
		return res{parts[1]}
	} else if parts[0] == "delete" {
		if len(parts) < 3 {
			res := new(commands.errorCommand)
			return res{"few arguments were introduced (must be 3 arguments)"}
		}
		res := new(commands.deleteCommand)
		return res{parts[1], parts[2]}
	}
	res := new(commands.errorCommand)
	return res{"incorrect command name"}
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
