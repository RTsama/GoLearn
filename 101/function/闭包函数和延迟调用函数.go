package main

import "fmt"

// 函数内部的函数可以访问 外面函数的局部变量 并将该函数返回。相对于内部函数而言，该变量可以一直访问，由编译器保持的一个全局变量
func bb() func(int2 int) {
	a := 1
	return func(int2 int) {
		//是函数bb回参的具体定义，与回参处声明一致
		a += int2
		fmt.Println(a)
	}
} //出参是个函数，该函数也可以有出参和入参
func later() {
	fmt.Println("defer修饰函数，使得函数被最后调用")
}

func main() {
	defer later()
	func1 := bb()
	func1(1)
	func1(1)
	func1(1)
	//闭包函数实现了对外面函数的局部变量的访问并保持 是编译器将外面函数的局部变量放入内存便于一直访问
}
