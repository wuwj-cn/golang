package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Println("HOME: ", HOME)
	fmt.Println("USER: ", USER)
	fmt.Println("GOROOT: ", GOROOT)
}
