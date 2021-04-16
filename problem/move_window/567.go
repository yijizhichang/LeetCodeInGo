package move_window

/*
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
 */

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	subL := len(s1)
	words := make([]uint32, 26)
	for _, this := range s1 {
		words[this-'a'] ++
	}
	tmp := make([]uint32, 26)
	match := 0
	copy(tmp, words)
	for left < len(s2) && left <= right{
		if match == subL {
			return true
		}
		if res:= tmp[s2[left]-'a'];res>0 {
			right++
			tmp[s2[left]-'a']--
			match++
		} else {
			left++
			right++
			copy(tmp, words)
			match = 0
		}
	}
	return false
}

// 官方
/*
由于排列不会改变字符串中每个字符的个数，所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的排列。
根据这一性质，记 s1的长度为 n，我们可以遍历 s2 中的每个长度为 n 的子串，判断子串和 s1中每个字符的个数是否相等，若相等则说明该子串是 s1 的一个排列。
使用两个数组 cnt1和 cnt2, cnt1 统计 s1 中各个字符的个数，cnt2 统计当前遍历的子串中各个字符的个数。
由于需要遍历的子串长度均为 n，我们可以使用一个固定长度为 n 的滑动窗口来维护 cnt2 滑动窗口每向右滑动一次，就多统计一次进入窗口的字符，少统计一次离开窗口的字符。然后，判断 cnt1 是否与 cnt2 相等，若相等则意味着 s1 的排列之一是 s2 的子串。
*/

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}