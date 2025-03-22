package commands

import (
	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	commandName string
	args        []string
}
