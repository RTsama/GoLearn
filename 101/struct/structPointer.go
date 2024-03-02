package main

import "fmt"

// 结构体声明 type 结构体名 struct {}
type Mystruct1 struct {
	name  string
	age   int
	sex   bool
	hobby []string
	home  //结构体嵌套结构体，可以不给结构体起成员名 直接用结构体名访问
}
type home struct {
	position string
}

// 结构体可以有自己的方法
// func (结构体名 结构体) 方法名
func (q *Mystruct1) song(name string) (ret string) { //传入结构体或结构体指针都可以
	ret = "NB"
	fmt.Printf("%v唱了一首%v,观众觉得%v", q.name, name, ret)
	q.home.position = "北京"
	P := q.home.position
	fmt.Println("他来自" + P)
	return ret
}
func functionTest() {
	qm := Mystruct1{
		name:  "唱歌的tom",
		age:   90,
		sex:   true,
		hobby: []string{"dance", "jump"},
	}
	re := qm.song("lll")
	fmt.Printf("\n%v", re)
}

//func pointer() {
//	var msp *Mystruct1
//	msp = &Mystruct1{
//		"Tom", 19, true, []string{"sing", "dance"},
//	}
//	msp.sex = false //结构体可以直接通过指针(内存地质)来操作
//	//*msp.age //此时会报错，因为相当于对msp.age取地址 .优先级更高
//	(*msp).name = "Changed" //使用括号不会报错
//	//go只可以对复杂数据类型 使用地址直接操作

//	fmt.Println(msp)
//}

func main() {
	//pointer()
	functionTest()
}
