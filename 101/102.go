package main

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2023-10-27"
	var url = "Code=%d&endDate=%s"
	//Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
	//Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	var target_url = fmt.Sprint(url, stockcode, enddate)

	fmt.Println(target_url)
	fmt.Println("今日是%s,密码是%d", enddate, stockcode)

}
