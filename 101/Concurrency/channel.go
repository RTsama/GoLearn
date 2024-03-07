package main

import (
	"fmt"
)

//channel是goroutine之间通讯的数据结构
//定义chan分为5种
//可读可取 c:=make(chan int)
//可读 var readChan <- chan int = c
//可取 var setChan chan<-int = c
//有缓冲 c:=make(chan int,5)
//无缓冲c:=make(chan int)

func channelTest1() {
	//c1 := make(chan int, 1)
	//c1 <- 1           //存入c1 (存chan)  若c1没缓冲区无法存入值
	//fmt.Println(<-c1) //读出数据(取chan) 不给缓冲区只有取数据时，才分配有容量 但此时没东西可以取，等待存入 死锁
	//c2 := make(chan int)
	c2 := make(chan int, 5)
	go func() { //使用自执行函数协程
		for i := 0; i < 10; i++ {
			c2 <- i
		}
	}()
	for j := 0; j < 10; j++ {
		fmt.Println(<-c2) //缺值阻塞时执行协程 获得c1
	}
}
func channelTest2() {
	c1 := make(chan int, 5)
	var readChan <-chan int = c1  //属于c1的一个只读channel
	var writeChan chan<- int = c1 //属于c1的一个只写channel
	writeChan <- 1
	fmt.Println(<-readChan)
	fmt.Println("--------")
	//close(c1) //关闭channel依旧可以取 但不能继续存
	writeChan <- 2
	writeChan <- 3
	writeChan <- 4
	writeChan <- 5
	writeChan <- 6 //可以通过属于c1的写channel 写c1
	close(c1)      //只有先close 才可以用for range
	for V := range c1 {
		fmt.Println(V)
	}
}
func channelTest3() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	c4 := make(chan int, 1)
	//channel的select方法 case后可执行 就会执行这个case
	c1 <- 1
	c2 <- 1
	c3 <- 1 //c1 c2 c3有值会随机执行或全不执行
	select {
	case <-c1:
		fmt.Println("channel1")
	case <-c2:
		fmt.Println("channel2")
	case <-c3:
		fmt.Println("channel3")
	case <-c4:
		fmt.Println("channel4")
	default:
		fmt.Println("No context")
	}
}
func main() {
	//channelTest1()
	//channelTest2()
	channelTest3()
}
