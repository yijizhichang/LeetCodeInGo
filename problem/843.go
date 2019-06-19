package problem
/*
我们给出了一个由一些独特的单词组成的单词列表，每个单词都是 6 个字母长，并且这个列表中的一个单词将被选作秘密。
你可以调用 master.guess(word) 来猜单词。你所猜的单词应当是存在于原列表并且由 6 个小写字母组成的类型字符串。
此函数将会返回一个整型数字，表示你的猜测与秘密单词的准确匹配（值和位置同时匹配）的数目。此外，如果你的猜测不在给定的单词列表中，它将返回 -1。
对于每个测试用例，你有 10 次机会来猜出这个单词。当所有调用都结束时，如果您对 master.guess 的调用不超过 10 次，并且至少有一次猜到秘密，那么您将通过该测试用例。
除了下面示例给出的测试用例外，还会有 5 个额外的测试用例，每个单词列表中将会有 100 个单词。
这些测试用例中的每个单词的字母都是从 'a' 到 'z' 中随机选取的，并且保证给定单词列表中的每个单词都是唯一的。

示例:
输入: secret = "acckzz", wordlist = ["acckzz","ccbazz","eiowzz","abcczz"]

 */
func cmp(sa, sb string) int {
	c := 0
	for i := 0; i < 6; i++ {
		if sa[i] == sb[i] {
			c++
		}
	}
	return c
}

func findSecretWord(wordlist []string, master *Master) {
	l := len(wordlist)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		t[i][i] = 6
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			c := cmp(wordlist[i], wordlist[j])
			t[i][j], t[j][i] = c, c
		}
	}

	m, j := l, 0
	for i := 0; i < l; i++ {
		cs, mc := [7]int{}, 0
		for _, j := range t[i] {
			cs[j]++
			if cs[j] > mc {
				mc = cs[j]
			}
		}
		if mc < m {
			m, j = mc, i
		}
	}

	wg := wordlist[j]
	res := master.Guess(wg)
	if 6 != res {
		sub := []string{}
		for _, w := range wordlist {
			if w != wg && cmp(w, wg) == res {
				sub = append(sub, w)
			}
		}
		findSecretWord(sub, master)
	}
}