package problem

import (
	"fmt"
	"sort"
)

/*给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"

示例 2:

输入: S = "aaab"
输出: ""

注意:

S 只包含小写字母并且长度在[1, 500]区间内。*/

/* 思路
1.出现次数最多的字符数不大于(n+1)/2  max<=(n+1)/2
2.提取k=>v,字符=>出现次数
3.每次从字符数排名前两位的字符中取字符，字符数相等时，避免，ab,ba,ab 情况出现
4.依次取出两个，直到取完为止*/

type CharNum struct {
	char string
	num  int
}

type CharNums []CharNum

//Len()
func (s CharNums) Len() int {
	return len(s)
}

//Less():排序
func (s CharNums) Less(i, j int) bool {
	return s[i].num > s[j].num
}

//Swap()
func (s CharNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getMaxNumChar(m map[string]int) (maxFirstChar, maxSecondChar string) {
	//重新生成二维数据结构
	var newStructs CharNums
	for k, v := range m {
		if v > 0 {
			newStructs = append(newStructs, CharNum{k, v})
		}
	}
	//倒序
	sort.Sort(newStructs)

	fmt.Println(newStructs)

	var v1, v2 int
	for k, v := range newStructs {
		if k == 0 {
			v1 = v.num
			maxFirstChar = v.char
		} else if k == 1 {
			v2 = v.num
			maxSecondChar = v.char
		} else {
			break
		}
	}

	//num 相当时，保持输出字符顺序不变
	if v1 == v2 {
		if maxSecondChar > maxFirstChar {
			maxFirstChar, maxSecondChar = maxSecondChar, maxFirstChar
		}
	}

	return
}

func reorganizeString(S string) string {

	m := make(map[string]int, 26)

	for i := 0; i < len(S); i++ {
		m[string(S[i])]++
	}

	maxChar, _ := getMaxNumChar(m)

	if m[maxChar] > (len(S)+1)/2 {
		return ""
	}

	var newStr string
	for j := 0; j < len(S)/2+1; j++ {
		maxFirstChar, maxSecondChar := getMaxNumChar(m)
		newStr += maxFirstChar
		newStr += maxSecondChar
		m[maxFirstChar]--
		m[maxSecondChar]--
	}

	return newStr
}
