package main

import (
	"fmt"
	"sync"
	"time"
)

var a string
var once sync.Once
var l sync.Mutex

func setup() {
	fmt.Println("set a value")
	a = "hello"
	//	l.Unlock()
}

func doPrint() {
	fmt.Println("enter doPrint")
	once.Do(setup)
	fmt.Println(a)
}
func main() {
	fmt.Println("run main")
	//	l.Lock()
	go doPrint()
	//	l.Lock()
	go doPrint()
	time.Sleep(5 * 1e9)
}
