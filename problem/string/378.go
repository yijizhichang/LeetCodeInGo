package string

/*
387. 字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。
*/

func firstUniqChar(s string) int {
	letter := make([]int32, 26)
	first := -1
	for _, this := range s {
		letter[this-'a']++
	}
	for idx, this := range s {
		if letter[this-'a'] == 1 {
			first = idx
			break
		}
	}
	return first
}
