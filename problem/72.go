package problem

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

*/
func minDistance(word1 string, word2 string) int {
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	word1Len, word2Len := len(word1), len(word2)
	if word1Len*word2Len == 0 {
		return word1Len + word2Len
	}

	dis := make([][]int, word1Len+1)
	for i := 0; i < word1Len+1; i++ {
		dis[i] = make([]int, word2Len+1)
	}
	for i := 0; i < word1Len+1; i++ {
		dis[i][0] = i
	}
	for j := 0; j < word2Len+1; j++ {
		dis[0][j] = j
	}

	for i := 1; i < word1Len+1; i++ {
		for j := 1; j < word2Len+1; j++ {
			if word1[i-1] == word2[j-1] {
				dis[i][j] = min(min(dis[i-1][j], dis[i][j-1]), dis[i-1][j-1]-1) + 1
			} else {
				dis[i][j] = min(min(dis[i-1][j], dis[i][j-1]), dis[i-1][j-1]) + 1
			}
		}
	}
	return dis[word1Len][word2Len]
}
