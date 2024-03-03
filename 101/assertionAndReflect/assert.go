package main

import "fmt"

// 断言 把一个接口类型指定为它原始类型
type user struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	class string
	user
}

func (u *user) sayMyName(name string) { //指针接收器
	fmt.Println("My name is", name)
}

//实现一个接口，必须实现这个接口提供的所有方法，
//定义一个方法，有值类型接收者和指针类型接收者两种。二者都可以调用方法
//因为 Go 语言编译器自动做了转换，所以值类型接收者和指针类型接收者是等价的。但
//是在接口的实现中，值类型接收者和指针类型接收者不一样
//!!!!!!
//以值类型接收者实现接口的时候，不管是类型本身，还是该类型的指针类型，都实现了该接口。
//以指针类型接收者实现接口的时候，只有对应的指针类型才被认为实现了该接口。

//有了接口和实现接口的类型，就会有类型断言。类型断言用来判断一个接口的值是否是实现该接口的某个具体类型。

func check(v interface{}) {
	//v.(*user).sayMyName(v.(*user).Name) //断言要满足方法要求，方法是一个指针接收器 v断言为*user
	switch v.(type) {
	//不同类型调用该方法时，用switch case控制。不如reflect简单
	case *user:
		fmt.Println("I'm a user")
	case *Student:
		fmt.Println("I'm a student")
	}
}

func assert1() {
	u := user{
		"Tom", 14, true,
	}
	s := &Student{
		"3.1",
		user{
			"Xiaoming", 18, true,
		},
	}
	check(&u)
	check(&s)
}

func main2() {
	assert1()
}
