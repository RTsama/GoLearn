package main

import "fmt"

// 结构体声明
type Mystruct struct {
	name  string
	age   int
	sex   bool
	hobby []string
}

// 显示声明
func initStruct() {
	var tom Mystruct
	tom.name = "Tom"
	tom.age = 18
	tom.sex = true
	tom.hobby = []string{"soccer", "basketball", "dance"}
	fmt.Println(tom)
}

// 隐式声明
func hiddenInit() {
	mike := Mystruct{
		"Mike", 18, true, []string{"sing", "dance", "rap"}, //顺序不能错
	}
	Nick := Mystruct{ //此时顺序可变
		name:  "Nick",
		sex:   false,
		hobby: []string{"sing", "dance"},
		age:   18,
	}
	fmt.Println(mike)
	fmt.Println(Nick)
}
func structFunction(mystruct Mystruct) {

}
func main1() {
	//initStruct()
	//hiddenInit()
	emptyStruct := new(Mystruct) //实例化一个空结构体
	fmt.Println(emptyStruct)

}
