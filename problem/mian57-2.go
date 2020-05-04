package problem

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

 

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
 

限制：

1 <= target <= 10^5

 */

func findContinuousSequence(target int) [][]int {
	i := 1
	j := 2
	res := make([][]int, 0)
	sum := 0
	for i > 0 && j < target && j > i{
		sum = (i + j)*(j-i+1)/2
		if sum > target {
			i++
		} else if sum < target {
			j++
		} else {
			tmp := make([]int, 0)
			for idx := i;idx <= j;idx++ {
				tmp = append(tmp, idx)
			}
			res = append(res, tmp)
			i++
		}
	}
	return res
}