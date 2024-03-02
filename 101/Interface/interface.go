package main

import "fmt"

// 接口是一系列方法的集合 实现了这些方法的结构属于这个接口
type animal interface {
	Eat()
	Run()
}
type Cat struct {
	Name string
	Sex  bool
}

//type Dog struct {
//	Name string
//}

// 此时Cat实现了Animal两个方法，已经实现了animal的接口
func (mio Cat) Run() {
	fmt.Printf("%v is running\n", mio.Name)
}

func (mio Cat) Eat() {
	fmt.Printf("%v is eatting\n", mio.Name)
}

func interface1() {
	//声明一个接口
	var a animal
	//声明一个变量
	c := Cat{
		Name: "Tom",
		Sex:  true,
	}
	a = c //c属于a接口 a中没有具体属性 但是具有方法  实现接口的方法的类 就可以用接口调用方法
	a.Run()
	a.Eat()
	//简写
	var b animal
	b = Cat{
		Name: "Jerry",
		Sex:  false,
	}
	b.Run()
	b.Eat()
}

// interface帮助实现了反省
func main() {
	interface1()
}
