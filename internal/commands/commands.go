package commands

import (
	"fmt"

	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	commandName string
	args        []string
}

type commands struct {
	allCommands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected 1 argument, the username\ncorrect usage: login [USERNAME]")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("The username has beed set successfully. Username is now %v", cmd.args[0])
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.allCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	err := c.allCommands[cmd.commandName](s, cmd)
	if err != nil {
		return err
	}
	return nil
}
