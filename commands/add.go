package commands

import (
	"strconv"
	"fmt"
	"github.com/Trafle/event_loop_realization/engine"
)

type addCommand struct {
	Arg1, Arg2 int
}

type printCommand struct {
	Arg string
}

func (a *addCommand) Execute(loop engine.Handler) {
	sum := a.Arg1 + a.Arg2
	loop.Post(&printCommand{Arg: strconv.Itoa(sum)})
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
