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
	a = c //c属于a接口 a中没有具体属性 但是具有方法  实现接口的方法的类 就可以用赋值接口调用方法
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

// interface帮助实现了范型
func myFunc(a interface{}) {
	//入参为空接口，所有东西都实现了无方法 所以函数可以接受任何参数
	fmt.Println(a)
}

func myFunc2(a animal) {
	a.Run()
	a.Eat()
}

var L animal

func myfunc3() {
	c := Cat{
		Name: "Tom",
		Sex:  false,
	}
	myInterface(c)
	L.Run() //想使L具有RUN方法 L是个全局变量接口
}

func myInterface(a animal) {
	L = a //传入一个animal接口的实例 传给L
}

func main() {
	//interface1()

	myFunc("hhh")
	a := 10
	myFunc(a)
	b := Cat{
		Name: "Jerry",
		Sex:  false,
	}
	myFunc(b)
	myFunc2(b) //该函数只接受满足animal接口的类型

	myfunc3()
}
