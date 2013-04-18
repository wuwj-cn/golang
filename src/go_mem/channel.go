package main

import (
	"fmt"
	"sync"
)

var c = make(chan int)
var c2 = make(chan int, 1)
var a string
var a2 string
var once sync.Once

func f() {
	a = "hello, world"
	<-c
}
func f2() {
	a2 = "hello, test"
	c2 <- 0
}

func setup() {
	fmt.Println("set a value")
	a = "hello"
}

func doPrint() {
	fmt.Println("enter doPrint")
	once.Do(setup)
	fmt.Println(a)
}
func main() {
	/*go f()
	c <- 0
	fmt.Println(a)
	go f2()
	<-c2
	fmt.Println(a2)*/
	fmt.Println("run main")
	go doPrint()
	go doPrint()
}
