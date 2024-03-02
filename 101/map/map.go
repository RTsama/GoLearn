package main

import "fmt"

func initMap() {
	var map1 map[string]string
	//var 变量名 map[key]value 这种声明方式需要赋值
	map1 = map[string]string{}
	map1["name"] = "Tom"
	map1["age"] = "134"
	map2 := map[string]string{}
	map2["name"] = "Tm"
	map2["age"] = "14"
	map3 := make(map[int]interface{}) //使用空接口接收任何数据类型
	map3[1] = true
	map3[12] = false
	map3[2] = "True"
	//删除
	delete(map3, 2)
	//使用string作为key的map可以使用for range便捷访问元素
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	fmt.Println(map1, map2, map3)

}

func main() {
	initMap()
}
