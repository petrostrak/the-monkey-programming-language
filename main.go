package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/petrostrak/the-monkey-programming-language/perl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	perl.Start(os.Stdin, os.Stdout)
}
