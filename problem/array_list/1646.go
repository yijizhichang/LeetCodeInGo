package array_list

/*
给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：

nums[0] = 0
nums[1] = 1
当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
返回生成数组 nums 中的 最大 值。

*/

func getMaximumGenerated(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[0], res[1] = 0, 1
	max := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + res[i/2+1]
		}
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}
