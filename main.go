package main

import (
	"fmt"
	"os"
	"os/user"

	"goethe/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the goethe programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
