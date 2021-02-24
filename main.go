package main

import (
	"fmt"
	"compilerDesign/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n",
		user.Username)
	fmt.Printf("Type commands here\n")
	repl.Start(os.Stdin, os.Stdout)
}

