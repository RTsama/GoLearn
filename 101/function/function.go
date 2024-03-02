package main

import "fmt"

func a(data1 int, data2 string) (ret1 int, ret2 string) {
	//出参位置写好了ret1 ret2 则不能在函数内重复定义ret1 ret2

	if data1 > 10 {
		fmt.Println(data1)
		return data1, "现在大于10"
	} else {
		fmt.Println(data2)
		return data1, "现在小于10"
	}

}
func b(data1 int, data2 int) (ret1 int, ret2 int) {
	//写好出参后 只写return意味着返回了ret1 ret2
	return
}
func mo(data1 int, data2 ...string) {
	//不定项参数...， 必须放在入参出参最后 此时不定项参数是一个切片(变长数组) 可用for range对其操作
	for k, v := range data2 {
		fmt.Println(k, v)
	}

	fmt.Println(data1, data2)
}

// func 函数名(入参)出参 {}

func main1() {
	r1, r2 := a(1, "Too small")
	fmt.Println(r1, r2)
	//go中 只有匿名函数可以在函数内部声明 编译器处理成在函数外声明函数 用变量来接收
	b := func(data1 string) {
		fmt.Println(data1)
	}
	b("Hello")
	mo(97123, "2", "3", "4")
	ar := []string{"1", "2", "3", "4"} //将数组作为不定项目参数全部传入函数 也要使用...
	mo(97872, ar...)
	//自执行函数(func(){})() 单纯自己执行一下
	(func() {
		fmt.Println("自执行函数的执行")
	})()

}
