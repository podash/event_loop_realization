package main

import(
	"github.com/Trafle/event_loop_realization/engine"
	"github.com/Trafle/event_loop_realization/commands"
	"os"
	"bufio"
)

func main () {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	
	inputFile := "example/example"
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
	 		cmd := parse(commandLine) // parse the line to get an instance of Command
	 		cmd.Execute()
	 	}
	}
	eventLoop.AwaitFinish()
}
