package main

import "fmt"

func ifElse() {
	a := 1
	if a > 1 {
		fmt.Println("a > 1")
	} else if a == 1 {
		fmt.Println("a = 1")
	} else {
		fmt.Println("a < 1")
	}
}

func switchCase() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
		fallthrough //使得继续往下一个条件走
	case 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("a is illegal")
	}
}

func forloop() {
	a := 0
	for a < 20 {
		a++
		if a == 5 {
			continue //跳过一轮
		}

		if a == 19 {
			break //跳过所有轮次
		}
		fmt.Println("a = ", a)

	}
	for b := 0; b < 10; b++ {
		fmt.Println("b = ", b)
	}

}

func gotoLang() {
	a := 0
A:
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 5 {
			goto B
		}
		if a == 10 {
			break A
		}

	}
B:
	for b := 20; b < 29; b++ {
		fmt.Println("b = ", b)
	}
}
func main() {
	//ifElse()
	//switchCase()
	//forloop()
	gotoLang()
}
