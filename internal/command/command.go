package command

import (
	"aggregator/internal/database"
	"aggregator/internal/state"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func HandlerLogin(s *state.State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("no argments")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("not enough arguments, username is required.")
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	return nil
}

func HandlerResetDb(s *state.State, cmd Command) error {
	err := s.Db.DeleteUsersInfo(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Users table reseted")
	return nil
}

func HandlerUsers(s *state.State, cmd Command) error {
	usersName, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, name := range usersName {
		if s.Cfg.CurrentUserName == name {
			name = name + " (current)"
		}
		println(name)
	}

	return nil
}

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no argments")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("not enough arguments, username is required.")
	}

	if _, err := s.Db.GetUser(context.Background(), cmd.Args[0]); err == nil {
		return fmt.Errorf("username already exists")
	}

	userparams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	createduser, err := s.Db.CreateUser(context.Background(), userparams)
	if err != nil {
		return err
	}

	logUserData(createduser)

	err = s.Cfg.SetUser(createduser.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	err := c.Handlers[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func logUserData(createduser database.User) {
	fmt.Printf("ID: %v \nCreatedAt: %v\nUpdatedAt: %v\nName: %v\n", createduser.ID, createduser.CreatedAt, createduser.UpdatedAt, createduser.Name)
}
