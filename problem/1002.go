package problem

/*
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。



示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]


提示：

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母

*/

/*
func commonChars(A []string) []string {
	var tmpNum [26]int
	for i := 0; i < 26; i++ {
		tmpNum[i] = -1
	}
	var finalNum [26]int
	for i := 0; i < 26; i++ {
		finalNum[i] = -1
	}
	for i := 0; i < len(A); i++ {
		var tempStr []byte = []byte(A[i])
		for j := 0; j < len(tempStr); j++ {
			if i == 0 {
				var tmp int = int(tempStr[j] - 'a')
				tmpNum[tmp]++
			}
			if i != 0 {
				var tmp int = int(tempStr[j] - 'a')
				finalNum[tmp]++
			}
		}
		if i > 0 {

			for k := 0; k < 26; k++ {
				if tmpNum[k] > -1 && finalNum[k] > -1 {
					if tmpNum[k] > finalNum[k] {
						tmpNum[k] = finalNum[k]

						//fmt.Println(k)
					}
				}else {
					tmpNum[k] = -1
				}
				finalNum[k] = -1
			}
		}
	}
	var ans []string
	for i := 0; i < 26; i++ {
		if tmpNum[i] != -1 {
			var str byte = byte(i + 'a')

			for j := 0; j <= tmpNum[i]; j++ {
				//fmt.Print(string(str), " ")
				ans = append(ans, string(str))

			}
		}
	}
	return ans
}

*/

func commonChars(A []string) []string {
	resM := make(map[string]int)
	for i, a := range A {
		temp := make(map[string]int)
		for _, i := range a {
			strI := string(i)
			if num, ok := temp[strI]; ok {
				temp[strI] = num + 1
			} else {
				temp[strI] = 1
			}
		}
		if i == 0 {
			for k, v := range temp {
				resM[k] = v
			}
		} else {
			for k, v := range temp {
				if ori, ok := resM[k]; ok {
					// 存在且小于原值替换
					if v < ori {
						resM[k] = v
					}
				} else {
					// 不存在则去除
					delete(resM, k)
				}
			}
			for k := range resM {
				if _, ok := temp[k]; !ok {
					// 不存在则去除
					delete(resM, k)
				}
			}
		}
	}
	res := make([]string, 0)
	for i, v := range resM {
		for v > 0 {
			res = append(res, i)
			v--
		}
	}
	return res
}
