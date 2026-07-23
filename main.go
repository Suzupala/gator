package main

import (
	"fmt"
	"os"

	"github.com/suzupala/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	currState := &state {
		config: &cfg,
	}
	commandList := &commands {
		cmds: make(map[string]func(*state, command) error),
	}
	commandList.register("login", handlerLogin) 
	
	if len(os.Args) <= 1 {
		fmt.Println("enter an argument")
		os.Exit(1)
	}
	commandName := os.Args[1]
	commandArguments := os.Args[2:]
	
	currCommand := command {
		name: commandName,
		args: commandArguments,
	}
	err = commandList.run(currState, currCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}




