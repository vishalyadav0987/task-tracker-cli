package cli

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}

func Parse(args []string) (*Command, error) {
	if len(args) < 2 {
		return nil, errors.New("no command provided")
	}

	cmd := &Command{
		Name: args[1],
	}

	if len(args) > 2 {
		cmd.Args = args[2:]
	}

	return cmd, nil
}
