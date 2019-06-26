package main

/*
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false

解题思路1：递归，分别取出s1和s2的每隔一个字符，如果s1或s2的最前面的字符和s3最前面的字符相同，则返回s1或s2的当前字符之后的字符串，再递归进行对比，直到s1和s2两个字符串的字符都为空，这个时候就返回true，否则返回false

阶梯思路2：动态规划，s1的前i个元素加上s2的前j个元素，等于s3的前i+j个元素，从题目来看假设从dp[0][0]的位置开始，走到dp[5][5]的位置，假如某个位置返回为真，那它的左边位置或者上边位置的值也同样为真，那它左边的位置就等于dp[i][j-1],和s3[i+j-1]的位置进行对比，如果为真，就是左边的路径为真，如果为假，那我们走上边的位置，等于dp[i-1][j],和[i+j-1]的位置进行对比，这两个对比，其中某一个为真，那最终的结果就是为真。看图

图中从[0][0]的位置开始找，与s3[0][0]进行对比，直到s[5][5]的位置

	   d b b c a
     0 1 2 3 4 5
	——————————————
  0| 1 0 0 0 0 0
a 1| 1 0 0 0 0 0
a 2| 1 1 1 1 1 0
b 3| 0 1 1 0 1 0
c 4| 0 0 1 1 1 1
c 5| 0 0 0 1 0 1
*/


func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	isInterleave(s1, s2, s3)
	isInterleave2(s1, s2, s3)

}
//动态规划
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)

	if len1+len2 != len3 {
		return false
	}
	s1b := []byte(s1)
	s2b := []byte(s2)
	s3b := []byte(s3)

	dp := make([][]bool, len(s2)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s1)+1)
	}

	for i := 0; i <= len2; i++ {
		for j := 0; j <= len1; j++ {

			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s1b[j-1] == s3b[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s2b[i-1] == s3b[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s2b[i-1] == s3b[i+j-1]) || (dp[i][j-1] && s1b[j-1] == s3b[i+j-1])
			}
		}
	}
	return dp[len(s2)][len(s1)]
}

//递归
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if (len(s1) + len(s2)) != len(s3) {
		return false
	}
	if s1 == "" && s2 == "" {
		return true
	} else if (s1 == "") {
		if s2[0] == s3[0] {
			return isInterleave(s1, s2[1:], s3[1:])
		}
	} else if (s2 == "") {
		if s1[0] == s3[0] {
			return isInterleave(s1[1:], s2, s3[1:])
		}
	}
	if (s1 != "" && s1[0] == s3[0] && s2[0] != s3[0]) {
		return isInterleave(s1[1:], s2, s3[1:])
	} else if (s2 != "" && s2[0] == s3[0] && s1[0] != s3[0]) {

		return isInterleave(s1, s2[1:], s3[1:])
	} else if (s1 != "" && s1[0] == s3[0] && s2 != "" && s2[0] == s3[0]) {

		isRel := isInterleave(s1[1:], s2, s3[1:])
		isRel2 := isInterleave(s1, s2[1:], s3[1:])
		if isRel || isRel2 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
