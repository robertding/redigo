package command

import (
	"errors"
	"fmt"
	"strings"
)

func Ping(_ string, _ []string) (string, error) {
	return "PONG", nil
}

func Echo(command string, args []string) (string, error) {
	if len(args) > 1 {
		return fmt.Sprintf("wrong number of arguments for '%s' command", command), errors.New("ERR")
	}
	return fmt.Sprintf("%s %s", command, strings.Join(args, " ")), nil
}
