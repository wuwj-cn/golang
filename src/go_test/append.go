package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	y := []int{7, 8, 9}
	x = append(x, y...)
	fmt.Println(x)
}
