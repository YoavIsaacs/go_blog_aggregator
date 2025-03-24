package main

import (
	"fmt"
	"os"

	"github.com/YoavIsaacs/go_blog_aggregator/internal/commands"
	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	state := commands.State{
		Config: c,
	}

	cmds := commands.Commands{
		AllCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("error: not enough arguments given")
		os.Exit(1)
	}

	cmd := commands.Command{
		CommandName: args[0],
		Args:        args[1:],
	}

	err = cmds.AllCommands[args[0]](&state, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
