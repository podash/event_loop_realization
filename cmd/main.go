package main

import(
	"github.com/Trafle/event_loop_realization/engine"
	"github.com/Trafle/event_loop_realization/commands"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

func parseAdd(arg1 string, arg2 string) engine.Command {
	int1, err := strconv.Atoi(arg1)
	int2, err := strconv.Atoi(arg2)
	if(err != nil) {return &commands.PrintCommand{Arg: err.Error()}}
	return &commands.AddCommand{Arg1: int1, Arg2: int2}
}

func parse(line string) engine.Command {
    parts := strings.Fields(line)
    cmd := parts[0]
    if(cmd == "print") {return &commands.PrintCommand{Arg: parts[1]}}
    if(cmd == "add") {return parseAdd(parts[1], parts[2])}
    if(cmd == "") {return &commands.PrintCommand{Arg: "No command specified"}}
    return &commands.PrintCommand{Arg: "TypeError: Unknown command"}
}


func main () {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	
	inputFile := "github.com/Trafle/event_loop_realization/cmd/example/example"
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
	 		cmd := parse(commandLine) // parse the line to get an instance of Command
	 		eventLoop.Post(cmd)
	 	}
	eventLoop.AwaitFinish()
}
