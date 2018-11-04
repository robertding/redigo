package command

import (
	"fmt"
	"strings"
)

func Ping(_ string, _ []string) string {
	return "PONG"
}

func Echo(command string, args []string) string {
	return fmt.Sprintf("%s %s", command, strings.Join(args, " "))
}
