package problem

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret := make([]int,len(nums))
	ret[0] = nums[0]
	max := ret[0]
	for i:=1;i<len(nums);i++ {
		tmp := ret[i-1] + nums[i]
		if tmp > nums[i] {
			ret[i] = tmp
		} else {
			ret[i] = nums[i]
		}
		if max < ret[i] {
			max = ret[i]
		}
	}
	return max
}
