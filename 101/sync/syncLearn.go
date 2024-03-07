package main

import (
	"fmt"
	"sync"
	"time"
)

//提供并法编程中的锁
//Mutex互斥锁 Lock() Unlock()
//RWMutex 读写互斥锁 Lock() Unlock() Rlock() Runlock()
//Once   Once.Do(一个函数) 这个方法无论被调用多少次 只会执行一次
//WaitGroup  Add(int) 设定需要Done多少次  wait在达到Done()执行次数之前一直阻塞
//Map 一个并法字典 Store Load LoadOrStore Range Delete
//Pool 并发池  Put Get
//Cond 没什么用的通知锁

func MutexLearn() {
	//只负责启动四个协程 不管协程是否执行完
	m := &sync.Mutex{} //保持对原值修改 传的是指针
	go mutexFunc(m)
	go mutexFunc(m)
	go mutexFunc(m)
	go mutexFunc(m)
	//保持协程之间相互独立
	//for {
	//}
}

func mutexFunc(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("+1")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
func RWMutex() {
	rwm := &sync.RWMutex{}
	go RmutexFunc(rwm)
	go RmutexFunc(rwm)
	go RmutexFunc(rwm)

	go RWmutexFunc(rwm)
	go RWmutexFunc(rwm)
	go RWmutexFunc(rwm)
	go RWmutexFunc(rwm)
	for {

	}

}
func RmutexFunc(lock *sync.RWMutex) {
	lock.RLock() //读锁 读取时不会互斥其他的读取协程 但是排斥写锁
	fmt.Println("Crazy")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}

func RWmutexFunc(lock *sync.RWMutex) {
	lock.Lock() //写入时排斥其他读锁和写锁
	fmt.Println("RW+1")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
func main1() { //go中 main函数也是个特殊的协程主协程  主协程结束所有子协程都会被强制终止 GMP
	//MutexLearn()
	RWMutex()
}
