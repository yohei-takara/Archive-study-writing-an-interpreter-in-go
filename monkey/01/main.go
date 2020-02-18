package main

import (
	"fmt"
	"os"
	"os/user"
	"yohei-takara/stary-writing-an-interpreter-in-go/monkey/01/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
