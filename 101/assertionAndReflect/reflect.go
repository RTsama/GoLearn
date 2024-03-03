package main

import (
	"fmt"
	"reflect"
)

// 在编译时不知道类型的情况下，可更新变量，运行时查看值、调用方法以及直接对他们的布局进行操作的机制，称为反射
// 可以获知数据的原始数据类型和数据内容/方法等，并可进行一定的操作
// reflect包 方法
// reflect.ValueOf() 获取输入参数接口中的数据的值
// reflect.TypeOf() 动态获取输入参数接口中的值的类型
// reflect.TypeOf().Kind() 用来判断类型
// reflect.ValueOf().Field(int) 用来获取值
// reflect.ValueOf().Elem() 获取原始数据并操作

type user1 struct {
	Name string
	Age  int
	Sex  bool
}
type student1 struct {
	Class string
	user1
}

func check1(inter interface{}) {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	fmt.Println("TypeOf:", t, "ValueOf", v)
	//t.NumField()//返回结构体中属性的个数
	for i := 0; i < t.NumField(); i++ {
		// reflect.ValueOf().Field(int) 用来获取值
		fmt.Println(v.Field(i))
	}
}
func check2(inter interface{}) {
	v := reflect.ValueOf(inter)
	fmt.Println(v)
	fmt.Println(v.FieldByIndex([]int{1, 0})) //第1层(从0开始索引)第0条
}
func reflectLearn() {
	s := student1{"3.1", user1{"BB", 14, false}}
	//u := user1{
	//	"AA", 21, true,
	//}
	//check1(u)
	//fmt.Println("----------------------------------")
	//check1(s)
	//fmt.Println("----------------------------------")
	//check2(s)
	//check3(s)
	//check3(u)
	check4(&s) //传地址保存更改
}
func check3(inter interface{}) {
	t := reflect.ValueOf(inter)
	ty := t.Kind()
	fmt.Println(ty)
	if ty == reflect.Struct {
		fmt.Println("STRUCT")
	}
	if ty == reflect.String {

	}

}
func check4(inter interface{}) {
	v := reflect.ValueOf(inter)
	// reflect.ValueOf().Elem() 获取原始数据并操作
	e := v.Elem()
	e.FieldByName("Class").SetString("4.2") //使用反射改变原始数据需要传入地址
	fmt.Println(inter)
}

func main() {
	reflectLearn()
}
