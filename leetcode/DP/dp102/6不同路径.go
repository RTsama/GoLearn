package main

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？

// 1确定dp数组（dp table）以及下标的含义 2确定递推公式 3dp数组如何初始化 4确定遍历顺序 5举例推导dp数组
// 二维DP 1 dp[i][j] 从【0,0】到[i][j]有dp[i][j]条路径 2 二维dp从两个方向推 dp[i][j] = dp[i-1][j]+dp[i][j-1]
// 3 初始化 dp[0][i]=1 dp[i][0]=1  4dp[i][j]都是从其上方和左方推导而来，那么从左到右一层一层遍历就可以了。

// 由于二维数组size是传入的，此为golang二维动态数组实现---切片
func uniquePaths(m int, n int) int {
	dp := make([][]int, m) //初始化m行 为nil的空切片
	for i := range dp {    //对dp的第一个维度进行遍历
		dp[i] = make([]int, n) //每一个维度创建一个长度为n的切片
		dp[i][0] = 1           //从起点可以直达的位置 直接赋成1
	}

	//var dp [m][n]int = [m][n]int{} //go中数组的初始化必须使用常量作初始化
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
