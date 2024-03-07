package main

import (
	"context"
	"fmt"
	"time"
)

//context包使用
//WithCancel 创建一个具有Clear()关闭函数的ctx
//WithDeadline 创建一个带有超时时间点的ctx
//WithTimeout 创建一个有超时时间段的ctx
//WithValue创建一个携带了参数的ctx
//Context基础接口 Deadline() Done() Err() Value()
// 返回空ctx TODO 不确定是否需要使用的时候使用  Backgroud 确定使用的时候使用

func son(flag chan bool, msg chan int) {
	t := time.Tick(time.Second) //创建一个通道，每秒向该通道发送一个时间事件
	//t := time.NewTicker()
	//t.C //结构体中自带的通道
	for _ = range t { //从通道中接收事件，使得循环每秒执行一次
		select {
		case m := <-msg:
			fmt.Printf("接收到值,%d\n", m)
		case <-flag:
			fmt.Println("Function is done")
			return
		}
	}
}

func goChannel() {
	flag := make(chan bool)
	message := make(chan int)
	go son(flag, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}
func ctxLearn() {
	//WithCancel 创建一个具有Clear()关闭函数的ctx  Backgroud() 确定使用ctx的时候使用返回一个空ctx
	ctx := context.WithValue(context.Background(), "name", "asd")
	ctx, clear := context.WithCancel(ctx) //此时具有key val属性
	//flag := make(chan bool)
	message := make(chan int)
	go son1(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	clear() //使上下文ctx关闭

}
func son1(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t { //从通道中接收事件，使得循环每秒执行一次
		select {
		case m := <-msg:
			fmt.Printf("接收到值,%d\n", m)
		case <-ctx.Done():
			fmt.Println("Function context is done", ctx.Value("name"))
			return
		}
	}

}

func main() {
	//goChannel()
	ctxLearn()
}
