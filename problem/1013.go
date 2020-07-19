package problem

/*
给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。



示例 1：

输出：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
示例 2：

输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false
示例 3：

输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

*/

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, num := range A {
		sum += num
	}

	if len(A) < 3 || sum%3 != 0 {
		return false
	}

	// 使用双指针'i','j', 尝试将数组A划分成 [0,i] (i,j) [j,len(A)-1]三部分
	var i, j int
	target := sum / 3
	temp1 := 0
	temp2 := 0
	for i = 0; i < len(A); i++ {
		temp1 += A[i]
		if temp1 == target {
			break
		}
	}

	for j = len(A) - 1; j >= 0; j-- {
		temp2 += A[j]
		if temp2 == target {
			break
		}
	}

	if i+1 < j {
		return true
	}
	return false
}
