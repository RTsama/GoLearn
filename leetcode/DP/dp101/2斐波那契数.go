package dp101

//F(0) = 0，F(1) = 1 F(n) = F(n - 1) + F(n - 2)，其中 n > 1 给你n ，请计算 F(n) 。

// 实战求法：1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
// 1) dp[i] 第i个 F(i) 2) dp[n] = dp[n-1] + dp[n-2] 3) dp[0] = 0, dp[1] = 1 4)从前到后遍历 5) 0 1 1 2 3 5 8 13 21 34
func fib0(n int) int {
	if n < 2 {
		return n
	}
	var dp [31]int
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, cur_sum := 0, 1, 0
	for i := 1; i < n; i++ { //1到n-1 计算n-2次 不算n=0 n=1
		cur_sum = a + b
		a = b
		b = cur_sum
	}
	return cur_sum
}
