package main

import (
	"fmt"
	"sync"
	"time"
)

func syncMap() {
	m := &sync.Map{}
	go func() {
		for {
			m.Store(1, 1)

		}
	}()
	go func() {
		for {
			fmt.Println(m.Load(1))
			m.Delete(1)
		}
	}()
	m.LoadOrStore(3, 3)

	m.LoadAndDelete(3)
	for {
	}
}
func SyncClass() {
	m := &sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(1 * time.Second)
		return true //返回true才会继续循环
	})
}

func main() {
	syncMap()
}
