package problem

/*

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。

示例 1：

输入: [1,3,5]
输出: 1
示例 2：

输入: [2,2,2,0,1]
输出: 0

*/

/*
题解

如下图所示，寻找旋转数组的最小元素即为寻找 右排序数组 的首个元素 numbers[x] ，称 xx 为 旋转点 。
排序数组的查找问题首先考虑使用 二分法 解决，其可将遍历法的 线性级别 时间复杂度降低至 对数级别 。

算法流程：
循环二分： 设置 ii, jj 指针分别指向 numbers 数组左右两端，m = (i + j) // 2m=(i+j)//2 为每次二分的中点（ "//" 代表向下取整除法，因此恒有 i \leq m < ji≤m<j ），可分为以下三种情况：
当 numbers[m] > numbers[j]时： mm 一定在 左排序数组 中，即旋转点 xx 一定在 [m + 1, j][m+1,j] 闭区间内，因此执行 i = m+1；
当 numbers[m] < numbers[j] 时： mm 一定在 右排序数组 中，即旋转点 xx 一定在[i, m][i,m] 闭区间内，因此执行 j = m；
当 numbers[m] == numbers[j] 时： 无法判断 mm 在哪个排序数组中，即无法判断旋转点 xx 在 [i, m][i,m] 还是 [m + 1, j][m+1,j] 区间中。解决方案： 执行 j = j - 1j=j−1 缩小判断范围 （分析见以下内容） 。
返回值： 当 i = j 时跳出二分循环，并返回 numbers[i] 即可。

*/

func findMin(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] > nums[j] {
			i = mid + 1
		} else if nums[mid] < nums[j] {
			j = mid
		} else {
			j--
		}
	}
	return nums[i]
}
