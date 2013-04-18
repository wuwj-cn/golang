package main

import "fmt"

func main() {
	//	fmt.Printf("hello world\n")
	//	var ar [3]int
	//	f(ar)
	//	fp(&ar)
	var aa = [10]int{2: 1, 5: 1, 7: 1}
	fpa(&aa)
	/*	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i*i, i*i*i})
	}*/
}
func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}
func fpa(a *[10]int) {
	fmt.Println(a)
}
