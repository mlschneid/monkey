package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mlschneid/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\nThis is the monkey programming language!\n", user.Username)
	fmt.Printf("Starting REPL...\n")
	repl.Start(os.Stdin, os.Stdout)
}
