package problem

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	i:=0
	for i<len(nums) {
		if nums[i] < target {
			i++
		} else if nums[i] == target {
			return i
		} else {
			break;
		}
	}
	return i
}
