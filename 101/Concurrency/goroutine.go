package main

import (
	"fmt"
	"sync"
	"time"
)

// 在调用一个方法的前面加上go就是goroutine 它会让方法异步执行 相当于协程(coroutine)
// 协程管理 var wg sync.WaitGroup 内部有一个计数器，然后围绕着计数器提供了三个方法
// wg.Add()给 WaitGroup 内部计数器的值增加
// delta wg.Done() 给 WaitGroup 内部计数器的值减一
// wg.Wait 如果计数器的值不为 0，那么此方法会阻塞，如果为 0，程序往下执行

func Run() {
	fmt.Println("I'm running")
}

func goTest() {
	go Run() //运行时机顺序与外面的goTest函数无关系
	i := 0
	for i < 200 {
		fmt.Println(i)
		i++
	}
	time.Sleep(1 * time.Second)
}
func Run1(wg *sync.WaitGroup) {
	fmt.Println("I'm running")
	wg.Done() //计数器-1 协程关闭
}
func syncTest() {
	var wg sync.WaitGroup
	wg.Add(1)    //说明有一个协程，计数器值为1
	go Run1(&wg) //指针传参 对同一个wg的计数器修改 否则死锁
	wg.Wait()    // 阻塞直到计数器值为0
}

func main1() {
	//goTest()
	syncTest()
}
