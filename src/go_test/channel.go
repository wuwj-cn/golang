package main
import ("fmt";"time";)

func sink(ch <-chan int){
	time.Sleep(5*1e9)
	for { fmt.Println(<-ch) }
	fmt.Println("received", <-ch)
}
func source(ch chan<- int){
	for { ch <- 1}
	fmt.Println("sent", 1)
}
func main(){
	c := make(chan int, 5)
//	go source(c)
//	go sink(c)
/*	go func(){
		time.Sleep(5*1e9)
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)*/
	go func(){
		for {
			time.Sleep(3*1e9)
			fmt.Print(<-c)
		}
	}()

	for {
		select {
//			case c <- 0:
//			case c <- 1:
			case v := <-c:
				fmt.Println("received", v)
			case <- time.After(3*1e9):
				fmt.Println("timed out after 3 seconds")
		}
	}
}
