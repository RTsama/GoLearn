package main

import (
	"fmt"
	"sync"
	"time"
)

func OnceFunc() {
	o := sync.Once{}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		o.Do(func() {
			fmt.Println("I'm in the Once")
		}) //同一个Once的 Do方法只有一个能执行
	}
}

func WGfunc() {
	wg := &sync.WaitGroup{}
	wg.Add(2) //还需done的次数
	go func() {
		time.Sleep(6 * time.Second)
		wg.Done()
		fmt.Println("sleep 6s")
	}()
	go func() {
		time.Sleep(8 * time.Second)
		wg.Done()
		fmt.Println("sleep 8s")
	}()
	//time.Sleep(5 * time.Second) //等待时间不够创建的协程执行 就结束了
	wg.Wait()
	fmt.Println("wake")
}

func main2() {
	//OnceFunc()
	WGfunc()
}
