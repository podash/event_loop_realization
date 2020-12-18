package main

import(
	"strings"
)

func Parse(line string) *func {
	parts := strings.Fields(line)
	cmd = parts[0]
	if(cmd == "print") {return &printCommand{Arg: parts[1]}}
	if(cmd == "add") {return &addCommand{Arg1: parts[1], Arg2: parts[2]}}
}