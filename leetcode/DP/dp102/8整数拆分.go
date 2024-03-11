package main

// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
// 1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
// tips 数值近似相等，乘积才尽可能大
// 1 dp[i] 对i进行拆分得到的最大乘积为dp[i] ; 2 dp[i] = j * dp[i-j] ; 3 dp[0] dp[1] 无意义 dp[2] = 1
func integerBreak(n int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max((j*(i-j)), j*dp[i-j]))
		}
	}
	return dp[n]
}
