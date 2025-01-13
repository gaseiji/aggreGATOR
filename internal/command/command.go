package command

import (
	"aggregator/internal/database"
	"aggregator/internal/rss"
	"aggregator/internal/state"
	"context"
	"errors"
	"fmt"
	"html"
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

func MiddlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
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

func HandlerFollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("incorrect number of arguments, expected 1")
	}

	totalFeeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	valid := false

	for _, v := range totalFeeds {
		if v.Url == cmd.Args[0] {
			valid = true
		}
	}

	if !valid {
		return errors.New("Url not found, please create a new feed using `addfedd` cmd or check mispelling Url")
	} else {

		curfeed, err := s.Db.GetFeed(context.Background(), cmd.Args[0])
		if err != nil {
			return err
		}

		newFeedFollowParams := database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    curfeed.ID,
		}

		followReturn, err := s.Db.CreateFeedFollow(context.Background(), newFeedFollowParams)
		logFollowData(followReturn)
	}

	return nil
}

func HandlerUnFollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Invalid number of arguments, expecting 1")
	}

	totalFeeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	valid := false

	for _, v := range totalFeeds {
		if v.Url == cmd.Args[0] {
			valid = true
		}
	}

	if !valid {
		return errors.New("Url not found, please create a new feed using `addfedd` cmd or check mispelling Url")
	} else {

		params := database.DeleteFeedFollowParams{
			Url:    cmd.Args[0],
			UserID: user.ID,
		}

		err := s.Db.DeleteFeedFollow(context.Background(), params)
		if err != nil {
			return err
		}

		fmt.Println("Unfollow executed sucessfully")

	}

	return nil
}

func HandlerFollowing(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("Invalid number of arguments, expecting none")
	}

	gotFeedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User: %v\n", user.Name)
	fmt.Printf("Following:\n")

	for _, v := range gotFeedFollows {

		fmt.Printf("%v\n", v.FeedNames)

	}

	return nil
}

func HandlerListFeeds(s *state.State, cmd Command) error {

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, v := range feeds {
		fmt.Printf("feed Name: %v\n", v.Name)
		fmt.Printf("feed Url: %v\n", v.Url)
		if v.UserName.Valid {
			fmt.Printf("feed UserName: %v\n", v.UserName.String)
		} else {
			fmt.Printf("feed UserName: NULL\n")
		}
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

func HanderAgg(s *state.State, cmd Command) error {
	newRssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	for _, v := range newRssFeed.Channel.Item {

		fmt.Printf("RssItem Title:")
		fmt.Println(html.UnescapeString(v.Title))
		fmt.Printf("RssItem Link:")
		fmt.Println(v.Link)
		fmt.Printf("RssItem Description:")
		fmt.Println(html.UnescapeString(v.Description))
		fmt.Printf("RssItem PubDate:")
		fmt.Println(v.PubDate[:16])
		fmt.Printf("\n\n")
	}
	fmt.Println()
	return nil
}

func HanderAddFeed(s *state.State, cmd Command, user database.User) error {
	//add a feed to the current logged user
	//check if there is 2 args
	if len(cmd.Args) != 2 {
		return errors.New("incorrect number of arguments, expected 2")
	}

	//set feed params
	feedparams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	//add to database
	newFeed, err := s.Db.CreateFeed(context.Background(), feedparams)
	if err != nil {
		return err
	}
	logFeedData(newFeed)

	newFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    newFeed.ID,
	}

	followReturn, err := s.Db.CreateFeedFollow(context.Background(), newFeedFollowParams)
	logFollowData(followReturn)

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

func logFeedData(data database.Feed) {
	fmt.Printf("ID: %v \nCreatedAt: %v\nUpdatedAt: %v\nName: %v\nUrl: %v\nUserID: %v\n", data.ID, data.CreatedAt, data.UpdatedAt, data.Name, data.Url, data.UserID)
}

func logFollowData(data database.CreateFeedFollowRow) {
	fmt.Printf("ID: %v \nCreatedAt: %v\nUpdatedAt: %v\nUserID: %v\nFeedID: %v\n", data.ID, data.CreatedAt, data.UpdatedAt, data.UserID, data.FeedID)
}
