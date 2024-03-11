package dp101

//数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
//每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
//请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯

// 1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
//1dp[i] 爬上第i台阶的最小花费 2）dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
//3dp[1] = cost[0] dp[2] = min(cost[0],cost[1]) no
//可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。” 也就是说 到达 第 0 个台阶是不花费的，但从 第0 个台阶 往上跳的话，需要花费 cost[0]。

func minCostClimbingStairs(cost []int) int {
	var dp [1000]int
	maxn := len(cost)
	if len(cost) <= 1 {
		return 0
	}
	//var 函数名 func(入参) 出参
	var min func(a, b int) int
	min = func(a, b int) int {
		if a >= b {
			return b
		} else {
			return a
		}
	}
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= maxn; i++ {
		dp[i] = min((dp[i-1] + cost[i-1]), (dp[i-2] + cost[i-2]))
	}
	return dp[maxn]
}
