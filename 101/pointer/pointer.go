package main

import "fmt"

func pointer1() {
	var a int
	a = 123
	fmt.Println(a)
	var b *int //声明一个int
	b = &a     //取地址依然与c语言相同
	*b = 456   //解引用
	fmt.Println(a)
	fmt.Println(a == *b)
}

func pointer2() {
	//数组指针和指针数组

	var arr [5]string
	arr = [...]string{"1", "2", "3", "4", "5"}
	//数组指针
	var arrP *[5]string
	arrP = &arr
	fmt.Println(arr, arrP)
	//指针数组
	var arrpp [5]*string //数组中元素类型是*string string类型的指针
	var str1 string = "1"
	var str2 string = "2"
	var str3 string = "3"
	var str4 string = "4"
	var str5 string = "5"

	arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}
	*arrpp[2] = "66"
	fmt.Println(*arrpp[2])
}

func pointerFunction(p1 *string) {
	*p1 = "CHANGE"
}
func main() {
	//pointer1()
	pointer2()

	str := "Hello"
	strp1 := &str
	//指针传参
	fmt.Println(str)
	pointerFunction(strp1)
	fmt.Println(str)
	//直接修改指针
	var str2 string = "测试地址"
	fmt.Println("str2:", str2)
	p := &str2 //内隐式定义一个指针
	*p = "通过指针更改了内容"
	fmt.Println("str2:", str2)
}
