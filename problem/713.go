package problem

func main() {
	arr := []int{10, 5, 2, 6}
	numSubarrayProductLessThanK(arr, 100)
}

/*
给定一个正整数数组 nums。

找出该数组内乘积小于 k 的连续的子数组的个数。

示例 1:

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
说明:

0 < nums.length <= 50000
0 < nums[i] < 1000
0 <= k < 10^6

解题思路：取两个数的乘积，若小于k，值付给临时乘积，再取下一个值乘以当前乘积，若大于k，记录个数，跳出循环，当前乘积除以left初始位置的参数，left+1，此时当前乘积就等于第二个，第三个参数的乘积，以此类推
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 0 {
		return 0
	}
	var count = 0         //小于k的个数
	var left = 0          //开始位置
	var right = 0         //结束位置
	var multiplier = 1    //临时乘积
	var lenth = len(nums) //数组长度
	//从头开始循环遍历数组
	for left = 0; left < lenth; left++ {

		for right < lenth && multiplier*nums[right] < k {

			multiplier *= nums[right]
			right++
		}
		if left == right {
			right++
		} else {
			count += right - left
			multiplier /= nums[left]
		}
	}
	return count

}
