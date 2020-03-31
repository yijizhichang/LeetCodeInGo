package problem

/*
给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。

 

示例 1：

输入：[1,2,3,4,4,3,2,1]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
示例 2：

输入：[1,1,1,2,2,2,3,3]
输出：false
解释：没有满足要求的分组。
示例 3：

输入：[1]
输出：false
解释：没有满足要求的分组。
示例 4：

输入：[1,1]
输出：true
解释：可行的分组是 [1,1]
示例 5：

输入：[1,1,2,2,2,2]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[2,2]

 */

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	num2Count := make(map[int]int)
	for _, cur := range deck {
		if _, ok := num2Count[cur];ok {
			num2Count[cur] ++
		} else {
			num2Count[cur] = 1
		}
	}
	final := 0
	flag := false
	for _,count := range num2Count {
		if flag {
			greater := gcd(final, count)
			if greater < 2{
				return false
			}
		}
		final = count
		flag = true
	}
	return true
}

// 最大公约数
func gcd(a,b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

