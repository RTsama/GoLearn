package main

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//机器人每次只能 向下 或者 向右 移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

// 1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
// 1 dp[i][j]从起点到达(i,j)的不同路径个数 2 dp[i][j]=dp[i-1][j]+dp[i][j-1] 如果(ij)是障碍点，则dp[i][j]为0
// 3 dp[i][0] = 1; dp[0][i] = 1 因为从(0, 0)的位置到(i, 0)的路径只有一条，所以dp[i][0]一定为1，dp[0][j]也同理。
// 但如果(i, 0) 这条边有了障碍之后，障碍之后（包括障碍）都是走不到的位置了，所以障碍之后的dp[i][0]应该还是初始值0。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	//起点终点有障碍直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	} //处于边界上障碍物后面的位置走不到
	//dp数组推导
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么dp[i][j]保持为0
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
