package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey language REPL!\n", user.Username)
	fmt.Println("Type .quit to exit the console.")

	repl.Start(os.Stdin, os.Stdout)
}
