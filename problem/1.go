package problem

/*
	O(n2)
 */
func twoSum1(nums []int, target int) []int {
	ret := make([]int, 0)
	if len(nums) < 2 {
		return ret
	}
	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<len(nums);j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i)
				ret = append(ret, j)
			}
		}
	}
	return ret
}

/*
  O(n)
 */
func twoSum2(nums []int, target int) []int {
	ret := make([]int,0)
	mp := make(map[int]int)
	if len(nums) < 2 {
		return ret
	}
	for i:=0;i<len(nums);i++ {
		if v,ok := mp[target-nums[i]];ok {
			ret = append(ret,i,v)
		} else {
			mp[nums[i]] = i
		}
	}
	return ret
}

