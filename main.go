package main

import (
	"./engine"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

type deleteCommand struct {
	arg1, arg2 string
}

func (d *deleteCommand) Execute(loop engine.Handler) {
	result := strings.Replace(d.arg1, d.arg2, "", -1)
	loop.Post(&printCommand{arg: result})
}

func parse(strCommand string) engine.Command {
	parts := strings.Fields(strCommand)

	if len(parts) == 0 {
		return &printCommand{"SYNTAX ERROR: empty command string"}
	}

	if parts[0] == "print" {
		if len(parts) < 2 {
			return &printCommand{"SYNTAX ERROR: few arguments were introduced (must be 2 arguments)"}
		}
		return &printCommand{parts[1]}
	} else if parts[0] == "delete" {
		if len(parts) < 3 {
			return &printCommand{"SYNTAX ERROR: few arguments were introduced (must be 3 arguments)"}
		}
		return &deleteCommand{parts[1], parts[2]}
	}
	return &printCommand{"SYNTAX ERROR: incorrect command name"}
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
