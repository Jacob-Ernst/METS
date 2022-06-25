package main

import (
	"errors"
	"fmt"
	"os"
)

// Inspired by https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go#using-flagset-to-implement-sub-commands
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
	Validate() error
}

func root(args []string) (err error) {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewCreateCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err = cmd.Init(os.Args[2:])
			if err != nil {
				return err
			}

			err = cmd.Validate()
			if err != nil {
				return err
			}

			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
