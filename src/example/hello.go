package main

import (
	"example/newmath"
	"fmt"
)

func main() {
	m := `hello
	world`
	fmt.Printf("m = %s\n", m) 
	fmt.Printf("hello world! Sqrt(2)=%v\n", newmath.Sqrt(2))
}
