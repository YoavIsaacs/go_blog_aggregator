package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/YoavIsaacs/go_blog_aggregator/internal/config"
	"github.com/YoavIsaacs/go_blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type State struct {
	Config config.Config
	DB     *database.Queries
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

	ctx := context.Background()

	_, err := s.DB.GetUser(ctx, cmd.Args[0])
	if err != nil {
		fmt.Println("error: user does not exist")
		os.Exit(1)
	}

	err = s.Config.SetUser(cmd.Args[0])
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

func HandleRegister(s *State, cmd Command) error {
	err := handleRegister(s, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return nil
}

func handleRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("error: expected 1 argument, username\ncorrect usage: login [USERNAME]")
	}

	ctx := context.Background()

	_, err := s.DB.GetUser(ctx, cmd.Args[0])
	if err == nil {
		fmt.Println("error: username already exists")
		os.Exit(1)
	}

	id := uuid.New()
	now := time.Now().UTC()

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.Args[0],
	}

	usr, err := s.DB.CreateUser(ctx, params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	s.Config.SetUser(usr.Name)

	fmt.Printf("user %s was created and set successfully\n", usr.Name)
	return nil
}

func HandleReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("error: expected no arguments")
	}

	ctx := context.Background()

	err := s.DB.ResetUsers(ctx)
	if err != nil {
		fmt.Println("error: error reseting users table")
		os.Exit(1)
	}

	fmt.Println("successfully reset users table")
	return nil
}

func HandleGetAllUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("error: expected no arguments")
	}

	ctx := context.Background()

	users, err := s.DB.GetAllUsers(ctx)
	if err != nil {
		fmt.Println("error: error getting all users from the database")
		os.Exit(1)
	}

	for _, user := range users {
		if s.Config.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func HandleAgg(s *State, cmd Command) error {
	rssFeed, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rssFeed)

	return nil
}
