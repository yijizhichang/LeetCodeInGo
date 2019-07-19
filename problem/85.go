package problem85

/*
	此题 为84题的升级版
*/
func maximalRectangle(mat [][]byte) int {
	m := len(mat)
	if m == 0 {
		return 0
	}

	n := len(mat[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = int(mat[0][j] - '0')
		for i := 1; i < m; i++ {
			if mat[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}
	max := 0
	for i := 0; i < m; i++ {
		tmp := largestRectangleArea(dp[i])
		if max < tmp {
			max = tmp
		}
	}

	return max
}

// 从 84 题复制
func largestRectangleArea(heights []int) int {
	ma := 0
	if len(heights) < 0 {
		return 0
	}

	var stack []int
	for i := 0; i < len(heights); i++ {
		if i == 0 {
			stack = append(stack, heights[i])
			m := maxArea(stack)
			ma = max(m, ma)
			continue
		}

		if heights[i-1] > heights[i] {
			m := maxArea(stack)
			ma = max(m, ma)
			stack = changeStack(stack, heights[i])
		}

		stack = append(stack, heights[i])

		if i+1 == len(heights) {
			m := maxArea(stack)
			ma = max(m, ma)
			continue
		}
	}

	return ma
}

func maxArea(h []int) int {
	ma := 0
	if len(h) == 1 {
		ma = h[0]
	}
	for i := 0; i < len(h); i++ {
		m := h[i] * (len(h) - i)
		ma = max(m, ma)
	}
	return ma
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func changeStack(s []int, insert int) []int {
	if len(s) == 0 {
		return append(s, insert)
	}
	for i := (len(s) - 1); i >= 0; i-- {
		if s[i] > insert {
			s[i] = insert
		}
	}

	return s
}
