package command

import (
	"errors"
	"fmt"
)

var DATA map[string]string

func init() {
	DATA = make(map[string]string)
}

func Set(command string, args []string) (string, error) {
	if len(args) != 2 {
		return fmt.Sprintf("wrong number of arguments for '%s' command", command), errors.New("ERR")
	}
	DATA[args[0]] = args[1]
	return "OK", nil
}

func Get(command string, args []string) (string, error) {
	if len(args) != 1 {
		return fmt.Sprintf("wrong number of arguments for '%s' command", command), errors.New("ERR")
	}
	return DATA[args[0]], nil
}
