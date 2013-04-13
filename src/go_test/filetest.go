package main

import (
	"./file"
	"fmt"
	"os"
)

func main() {
	hello := []byte("hello workd\n")
	file.Stdout.Write(hello)
	file, err := file.Open("/does/not/exist", 0, 0)
	if file == nil {
		fmt.Printf("can't open file; err=%s\n", err)
		os.Exit(1)
	}
}
