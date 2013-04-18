package main

import "fmt"

type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float64
}
type C struct {
	x float64
	int
	string
}
type D struct {
	ax, ay, az int
}
type E struct {
	A
	D
}

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("haha\n")
	//fmt.Printf(test()+"\n")
	//	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var sl = make([]int, 0, 5)
	fmt.Println(sl)
	sl = appendToSlice(11, sl)
	fmt.Println(sl)
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	c := C{3.5, 5, "hello"}
	fmt.Println(c.x, c.int, c.string)
	e := E{A{1, 2}, D{33, 4, 5}}
	fmt.Println(e.A.ax, e.D.ax)
}
func test() int {
	return 2 * 2
}
func appendToSlice(i int, sl []int) []int {
	if len(sl) == cap(sl) {
		fmt.Println("error")
	}
	n := len(sl)
	sl = sl[0 : n+1]
	sl[n] = i
	return sl
}
