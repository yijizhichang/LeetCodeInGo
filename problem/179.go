package problem

/*
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

func largestNumber(nums []int) string {

	arrLen := len(nums) //空数组的情况下
	if arrLen == 0 {
		return "0"
	}

	sum := 0 //数组参数都是0的情况下
	for _, num := range nums {
		sum += num
	}
	if sum == 0 {
		return "0"
	}

	length := len(nums)
	for i := 0; i < length; i++ {

		for j := 0; j < length-i-1; j++ {

			s1 := strconv.Itoa(nums[j]) + strconv.Itoa(nums[j+1])
			s2 := strconv.Itoa(nums[j+1]) + strconv.Itoa(nums[j])
			s1Int, _ := strconv.Atoi(s1)
			s2Int, _ := strconv.Atoi(s2)

			if s1Int < s2Int {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	s := make([]string, length)
	for i := range nums {
		s[i] = strconv.Itoa(nums[i])
	}
	return strings.Join(s, "")

}
