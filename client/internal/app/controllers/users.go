package ctr

import (
	"chat_cli/internal/app/services"
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

const (
	RoleUser  int32 = 0
	RoleAdmin int32 = 1
)

type UserController struct {
	s *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{s}
}

func (c *UserController) Create(cCtx *cli.Context) error {
	name := cCtx.Args().Get(0)
	if name == "" {
		return errors.New("name (arg[0]) is required")
	}

	email := cCtx.Args().Get(1)
	if email == "" {
		return errors.New("email (arg[1]) is required")
	}

	password := cCtx.Args().Get(2)
	if password == "" {
		return errors.New("password (arg[2]) is required")
	}

	roleID := RoleUser
	if cCtx.String("role") == "admin" {
		roleID = RoleAdmin
	}

	id, err := c.s.Create(cCtx.Context, name, email, password, int32(roleID))
	if err != nil {
		return err
	}

	fmt.Printf("user with id = %v has been created\n", id)
	return nil
}

func (c *UserController) Get(cCtx *cli.Context) error {
	rawID := cCtx.Args().Get(0)
	if rawID == "" {
		return errors.New("id (arg[0]) is required")
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("failed to convert id (%v) to int: %v", id, err)
	}

	user, err := c.s.GetByID(cCtx.Context, int64(id))
	if err != nil {
		return err
	}

	fmt.Println(user.String())
	return nil
}

func (c *UserController) List(cCtx *cli.Context) error {
	users, err := c.s.List(cCtx.Context)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("Seems like no users haven't been added yet")
		return nil
	}

	fmt.Println("Users:")
	for _, u := range users {
		fmt.Println(u.String())
	}
	return nil
}

func (c *UserController) Update(cCtx *cli.Context) error {
	rawID := cCtx.Args().Get(0)
	if rawID == "" {
		return errors.New("id (arg[0]) is required")
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("failed to convert id (%v) to int: %v", id, err)
	}

	email := cCtx.String("email")
	name := cCtx.String("name")
	if err := c.s.Update(cCtx.Context, int64(id), email, name); err != nil {
		return err
	}

	fmt.Printf("User(id=%v) has been updated\n", id)
	return nil
}

func (c *UserController) Delete(cCtx *cli.Context) error {
	rawID := cCtx.Args().Get(0)
	if rawID == "" {
		return errors.New("id (arg[0]) is required")
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("failed to convert id (%v) to int: %v", id, err)
	}

	if err := c.s.Delete(cCtx.Context, int64(id)); err != nil {
		return err
	}

	fmt.Printf("User(id=%v) has been deleted\n", id)
	return nil
}
