package main

import (
	"fmt"
	"strconv"
)

func nonimate() {
	var v0 int
	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4
	fmt.Println(v0, v1, v2, v3, v4)
}

func str2int() {

	string1 := "123"

	int1, _ := strconv.Atoi(string1)
	fmt.Printf("%T to %T", string1, int1)
}

func main() {
	str2int()
}
