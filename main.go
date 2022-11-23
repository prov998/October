package main

import (
	"October/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello! This is the october language")
	repl.Start(os.Stdin, os.Stdout)
}
