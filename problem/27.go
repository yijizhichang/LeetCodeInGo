package problem

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for ;index<len(nums); {
		if nums[index] == val {
			nums = append(nums[:index],nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}
