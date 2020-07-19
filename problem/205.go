package problem

/*
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

输入: s = "egg", t = "add"
输出: true
示例 2:

输入: s = "foo", t = "bar"
输出: false
示例 3:

输入: s = "paper", t = "title"
输出: true
说明:
你可以假设 s 和 t 具有相同的长度。
*/

func isIsomorphic(s string, t string) bool {
	s2t := make(map[string]string)
	for index, this := range s {
		if value, ok := s2t[string(this)]; ok {
			if value != string(t[index]) {
				return false
			}
		} else {
			s2t[string(this)] = string(t[index])
		}
	}
	t2s := make(map[string]string)
	for index, this := range t {
		if value, ok := t2s[string(this)]; ok {
			if value != string(s[index]) {
				return false
			}
		} else {
			t2s[string(this)] = string(s[index])
		}
	}
	return true
}

/*
记录个数
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
}
*/
