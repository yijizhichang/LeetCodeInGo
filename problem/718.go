package problem

/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。



示例：

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。


提示：

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100

*/

/*
解法： 滑动窗口
想象两把尺子，交错滑动第二把尺子，找出最长段即可
*/

func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	ret := 0
	// B不动A动
	for i := 0; i < n; i++ {
		length := min(m, n-i)
		ret = max(ret, maxLength(A, B, i, 0, length))
	}
	// A不动B动
	for i := 0; i < m; i++ {
		length := min(n, m-i)
		ret = max(ret, maxLength(A, B, 0, i, length))
	}
	return ret
}

func maxLength(A, B []int, addA, addB, length int) int {
	ret, k := 0, 0
	for i := 0; i < length; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}
