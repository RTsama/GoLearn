package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type person struct {
	name string
	age  uint
	addr address
}
type address struct {
	province string
	city     string
}

func (p *person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// 通过工厂函数创建自定义结构体的方式，可以让调用者不用太关注结构体内部的字段，
// 只需要给工厂函数传参就可以了。
// 用下面的代码，即可创建一个*person 类型的变量 p1：

func NewPerson(name string) *person {
	return &person{name: name}
}
func main1() {
	var s fmt.Stringer
	p1 := NewPerson("ZhangSan")
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
	//为 s 的确是一个 *person，所以不会异常，可以正常返回 p2。但是如果我再添加如下代码
	//，对 s 进行 address 类型断言，就会出现一些问题：
	//panic: interface conversion: fmt.Stringer is *main.person, not main.address
	a := s.(address)
	fmt.Println(a)
}
