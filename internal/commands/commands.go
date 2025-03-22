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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected 1 arguments, the username\ncorrect usage: login [USERNAME]")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("The username has beed set successfully. Username is now %v", cmd.args[0])
	return nil
}
