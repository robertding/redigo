package command

import "fmt"

func Default(command string, _ []string) string {
	return fmt.Sprintf("Hello I`m Redigo and I Dont understand %s", command)
}
