package problem

/*
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，chars 中的每个字母都只能用一次。

返回词汇表 words 中你掌握的所有单词的 长度之和。

 

示例 1：

输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6
解释：
可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
示例 2：

输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
输出：10
解释：
可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。

 */

func countCharacters(words []string, chars string) int {
	letter2Count := make(map[rune]int)
	for _, char := range chars {
		if count, ok := letter2Count[char]; ok {
			letter2Count[char] = count + 1
		} else {
			letter2Count[char] = 1
		}
	}
	res := 0
	for _, word := range words {
		thisLetter2Count := make(map[rune]int)
		for key, value := range letter2Count {
			thisLetter2Count[key] = value
		}
		thisCanSpell := true
		for _, w := range word {
			if targetCount, ok := thisLetter2Count[w]; !ok || targetCount == 0 {
				thisCanSpell = false
				break
			} else {
				thisLetter2Count[w] = targetCount - 1
			}
		}
		if thisCanSpell {
			res += len(word)
		}
	}
	return res
}

/*
func countCharacters(words []string, chars string) int {
    var byteCount [26]int
    for _, char := range chars {
        byteCount[char-'a']++
    }

    ret := 0
    for _, word := range words {
        bc, match := byteCount, true
        for _, char := range word {
            if bc[char-'a'] <= 0 {
                match = false
                break
            }
            bc[char-'a']--
        }
        if match {
            ret += len(word)
        }
    }
    return ret
}
 */
