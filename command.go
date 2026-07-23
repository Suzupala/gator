package main

import (
	"fmt"
)

type commands struct {
	cmds map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

func (c *commands) run(s *state, cmd command) error {
	call, ok := c.cmds[cmd.name]	
	if !ok {
		return fmt.Errorf("command does not exist")
	}
	return call(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

