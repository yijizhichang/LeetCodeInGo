package problem

/*
	给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

	示例 1:

	输入: nums = [1,1,1,2,2,3], k = 2
	输出: [1,2]
	示例 2:

	输入: nums = [1], k = 1
	输出: [1]
	说明：

	你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
	你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/

func topKFrequent(nums []int, k int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	maxCount := 0
	element2CountMap := make(map[int]int, 0)
	for _,num := range nums {
		element2CountMap[num]++
		if element2CountMap[num] > maxCount {
			maxCount = element2CountMap[num]
		}
	}
	tmp := make([][]int,maxCount+1)
	for k,v := range element2CountMap{
		tmp[v] = append(tmp[v], k)
	}
	for i:=maxCount;i>=0;i--{
		if len(tmp[i]) == 0 {
			continue
		}
		ret = append(ret, tmp[i]...)
		if len(ret) == k {
			break
		}
	}
	return ret
}