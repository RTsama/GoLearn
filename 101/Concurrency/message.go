package main

import "fmt"

func setChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("I'm in setChan")
		writec <- i
	}
}

func getChan(read <-chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println("I'm in getChan, get", <-read, "from setChan")
	}
}

func test1() {
	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c
	go setChan(writec)

	getChan(readc)
}

func main() {
	test1()
}
