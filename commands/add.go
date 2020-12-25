package commands

import (
	"strconv"
	"fmt"
	"github.com/Trafle/event_loop_realization/engine"
)

type AddCommand struct {
	Arg1, Arg2 int
}

type PrintCommand struct {
	Arg string
}

func (a *AddCommand) Execute(loop engine.Handler) {
	sum := a.Arg1 + a.Arg2
	loop.Post(&PrintCommand{Arg: strconv.Itoa(sum)})
}

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
