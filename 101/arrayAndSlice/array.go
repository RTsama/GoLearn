package main

import (
	"fmt"
	"sort"
)

const (
	c1 byte = 65
	c2      = 'a'
	c3 rune = '中'
)

func test01() {
	fmt.Printf("c1的码值=%v, 字符为%c， 类型是%T\n", c1, c1, c1)
	fmt.Printf("c1的码值=%v, 字符为%c， 类型是%T\n", c2, c2, c2)
	fmt.Printf("c1的码值=%v, 字符为%c， 类型是%T\n", c3, c3, c3)

}
func array1() {
	//[元素长度]元素类型{元素值}
	var nameList [3]string = [3]string{"Mike", "Tom", "Tony"}
	fmt.Println(nameList)
	fmt.Printf("%v,的类型是：%T\n", nameList[1], nameList[1])
	b := [...]int{1, 3, 23, 3, 1, 123, 2, 123, 123} //不确定个数时可以用...填充
	fmt.Println(b)
	var c = new([15]string)
	c[10] = "tom"
	fmt.Println(c)
	zoom := [...]string{"cat", "dog", "eagle"}
	for i := 0; i < len(zoom); i++ {
		fmt.Println(zoom[i], "run")
		fmt.Println(zoom[i] + "run") //用加号连接没空格
	}
	//for range循环方式
	for i, v := range zoom {
		fmt.Println(i, v)
	}
	fmt.Println("zoom capacitiy", cap(zoom))
	fmt.Println("zoom length", len(zoom))
	//len() 长度  cap()容量 数组条件下 两者大小相同
	erArray := [3][3]int{
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
	}
	fmt.Println(erArray)
}
func slice1() {
	//声明后长度可变,切片内必须为同种元素
	var nameList []string
	fmt.Println(nameList == nil)
	var nameList1 = make([]string, 2)
	fmt.Println(nameList1)
	nameList = append(nameList, "Tom")
	nameList = append(nameList, "Jerry")
	nameList = append(nameList, "Mike")
	ageList := make([]int, 3)
	fmt.Println(ageList)
	array := [3]int{1, 2, 3}
	slices := array[:] //从头 切到尾
	//切片后可以使用append改变长度
	fmt.Println(slices)
	fmt.Println(array[0:2]) //左闭右开
	fmt.Println(array[:2])  //左闭右开 从头和到尾可以省略
	fmt.Println(array[1:2]) //左闭右开

}
func slice2() {
	//切片排序
	var ints = []int{4, 3, 5, 6}
	sort.Ints(ints)
	fmt.Println("升序", ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("降序", ints)
}

func array2slice() {
	a := [3]int{0, 1, 2}
	cl := a[1:]
	cl = append(cl, 5)
	//只有切片才能改变大小 append
	cl[0] = 5
	//对切片的修改不会影响原数组
	fmt.Println("a = ", a, " cl = ", cl)
	fmt.Println("cl len", len(cl), " cl cap", cap(cl))
	//此时切片的长度为3 容量为4 随着切片变大超过cap时候 go会自动为cap增大到另一个值
	fmt.Println(cl)
	cl1 := a[2:]
	copy(cl, cl1) //将cl1对应位置元素覆盖到cl上
	fmt.Println(cl, cl1)
	copy(cl[2:], cl1) //参数是两个切片
	fmt.Println(cl, cl1)
	//不赋值创造空切片
	var aa []int
	fmt.Printf("a's type is %T\n", a)
	fmt.Printf("aa's type is %T\n", aa) //此时不赋初值是切片 空切片
	//使用make创造切片
	aaa := make([]int, 4, 6)              //len 4  cap 6   cap可以省略
	fmt.Printf("aaa's type is %T\n", aaa) //此时不赋初值是切片 空切片
}

func main() {
	//test01()
	//array1()
	//slice1()
	//slice2()
	array2slice()
}
