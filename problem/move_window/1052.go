package move_window

/*
今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。

在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。

书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。

请你返回这一天营业下来，最多有多少客户能够感到满意的数量。


示例：

输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16
解释：
书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.


提示：

1 <= X <= customers.length == grumpy.length <= 20000
0 <= customers[i] <= 1000
0 <= grumpy[i] <= 1

*/

func maxSatisfied(customers []int, grumpy []int, X int) int {
	total := 0
	n := len(customers)
	// 算下原本不生气时间段时可以获得的满意度
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	// 从0开始保持一个X大小的窗口，作为初始本身就获取不到的值，通过特殊手段才能额外获取的
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := X; i < n; i++ {
		// 去头加尾
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
