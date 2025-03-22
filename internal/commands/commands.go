package commands

import (
	"fmt"

	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
)

type State struct {
	Config config.Config
}

type Command struct {
	CommandName string
	Args        []string
}

type Commands struct {
	AllCommands map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("expected 1 argument, username\ncorrect usage: login [USERNAME]")
	}

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("The username has been set successfully. Username is now %v", cmd.Args[0])
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.AllCommands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	err := c.AllCommands[cmd.CommandName](s, cmd)
	if err != nil {
		return err
	}
	return nil
}
