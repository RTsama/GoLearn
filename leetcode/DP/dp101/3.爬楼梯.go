package dp101

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。

// 1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
// 1.dp[i]上到第i阶的方法数 2. dp[n] = d[n-1] + (dp[n-2]) n-1上一阶 n-2 上两阶  3.dp[1] = 1 dp[2] = 2. 4.从前到后 5. 1 2 3 5 8 13
// 注意 n-2 虽然还剩两级 可以1+1 2 但是n-2上一阶是n-1的解  思考方式是dp[i]如何由前序状态获得
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var dp [46]int
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	a := 1
	b := 2
	cur := 0
	for i := 0; i < (n - 2); i++ {
		cur = a + b
		a = b
		b = cur
	}
	return cur
}
