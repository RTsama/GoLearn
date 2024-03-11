package main

// 题意：给定两个数组，编写一个函数来计算它们的交集。
func intersection(num1 []int, num2 []int) []int {
	set := make(map[int]struct{}, 0) //用map模拟set
	res := make([]int, 0)
	for _, v := range num1 {
		// index,value  range中的两项
		if _, ok := set[v]; !ok { //set中 有这个key 则返回键值和true ; 没有这个key，返回nil和false
			set[v] = struct{}{}
		}
	}
	for _, v := range num2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
