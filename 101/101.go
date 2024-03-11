package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct{}

func (d Dog) Sound() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Sound() {
	fmt.Println("Meow!")
}

func main() {
	//animals := []Animal{Dog{}, Cat{}}
	//
	//for _, animal := range animals {
	//	// 判断接口类型是否是Dog
	//	if dog, ok := animal.(Dog); ok {
	//		dog.Sound()
	//	}
	//
	//	// 判断接口类型是否是Cat
	//	if cat, ok := animal.(Cat); ok {
	//		cat.Sound()
	//	}
	//}
	dp := make([][]int, 5) //初始化m行 为nil的空切片
	for i := range dp {    //对dp的第一个维度进行遍历
		dp[i] = make([]int, 6) //每一个维度创建一个长度为n的切片
		dp[i][0] = 1           //从起点可以直达的位置 直接赋成1
	}
	for i, j := range dp {
		println("i=", i, "j=", j)
	}

}
