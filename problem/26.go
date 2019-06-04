package problem

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	left, right := 0, 1
	for ;right<len(nums);right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	nums = nums[:left+1]
	return len(nums)
}
