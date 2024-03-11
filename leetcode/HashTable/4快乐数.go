package main

//快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，
//然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
//如果 n 是快乐数就返回 True ；不是，则返回 False 。

//思路：模拟并记住中间过程的值，如果后面出现曾出现的中间值，说明进入死循环

func isHappy(n int) bool {

	var getSum func(int) int
	getSum = func(m int) int {
		sum := 0
		for m > 0 {
			a := m % 10
			sum += a * a
			m = m / 10
		}
		return sum
	}
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true //中间结果存入哈希表
	} // 已存在中间变量直接跳出死循环
	return n == 1 //n不是1 说明不是快乐数
}
