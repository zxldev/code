package code

/**
@see https://leetcode-cn.com/problems/distinct-subsequences/
*/
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	/**
		dp矩阵
		t\s "" b a b g b a g (m,j)
		""  1  1 1 1 1 1 1 1
	 	b   0  1 1 2 2 3 3 3
		a   0  0 1 1 1 1 4 4
		g   0  0 0 0 1 1 1 5
		(n,i)
	*/

	/**

	考虑2种情况
	情况1：
		结尾不同
		示例：  s: bag  t:ba
		这时候无论结尾是什么例如  bax bay baz ,都不能改变结果， 都和  s: ba  t: ba 结果一致
		所以有：  dp[i][j] = dp[i-1][j]
	情况2：
	   结尾相同
		示例：  s:babg t: bag
		这时候有两种情况
		   情况2.1：s末尾参与计算 ，两个末尾的g对应上了 ， 则等效于 s:bab t:ba 的结果 ，既：dp[i-1][j-1]
		   情况2.2：s末尾不参与计算 ，则等效于 s:bab  t:bag 的结果 既： dp[i-1][j]
		所以有：   dp[i][j] = dp[i-1][j-1]+dp[i-1][j]
	*/

	//template t
	dp := make([][]int, n+1)
	//init dp
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = 0
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s[j-1] != t[i-1] {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]

}
