package move_window

/*
1423. 可获得的最大点数
几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。

每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。

你的点数就是你拿到手中的所有卡牌的点数之和。

给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。



示例 1：

输入：cardPoints = [1,2,3,4,5,6,1], k = 3
输出：12
解释：第一次行动，不管拿哪张牌，你的点数总是 1 。但是，先拿最右边的卡牌将会最大化你的可获得点数。最优策略是拿右边的三张牌，最终点数为 1 + 6 + 5 = 12 。
示例 2：

输入：cardPoints = [2,2,2], k = 2
输出：4
解释：无论你拿起哪两张卡牌，可获得的点数总是 4 。
*/

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	minL := 0
	for _, this := range cardPoints[:n-k] {
		minL += this
	}
	tmp := minL
	for i := n - k; i < n; i++ {
		// 滑动窗口每向右移动一格，增加从右侧进入窗口的元素值，并减少从左侧离开窗口的元素值
		tmp = tmp - cardPoints[i-(n-k)] + cardPoints[i]
		minL = min(minL, tmp)
	}
	total := 0
	for _, this := range cardPoints {
		total += this
	}
	return total - minL
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
