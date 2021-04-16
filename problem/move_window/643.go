package move_window
/*
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

 

示例：

输入：[1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

 */

func findMaxAverage(nums []int, k int) float64 {
	maxN := 0
	for i:=0; i<k; i ++ {
		maxN += nums[i]
	}
	tmp := maxN
	for i:=k; i<len(nums); i++ {
		tmp = tmp + nums[i] - nums[i-k]
		if tmp > maxN {
			maxN = tmp
		}
	}
	return float64(maxN)/float64(k)
}
