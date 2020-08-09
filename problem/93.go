package problem

import "strconv"

/*
93. 复原IP地址
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]


链接：https://leetcode-cn.com/problems/restore-ip-addresses

*/
func restoreIpAddresses(s string) []string {
	var ipAddress []string
	//先判断字符串长度是否合法
	if len(s) < 4 {
		return ipAddress
	}
	//循环校验IP地址是否合法
	for x := 1; x <= 3 && x <= len(s)-3; x++ {
		if !checkIP(s[:x]) {
			continue
		}
		//第二位数
		for y := 1; y <= 3 && y <= len(s)-x-2; y++ {
			if !checkIP(s[x : x+y]) {
				continue
			}
			//第三位数
			for z := 1; z <= 3 && z <= len(s)-x-y-1; z++ {
				if !checkIP(s[x+y:x+y+z]) || !checkIP(s[x+y+z:]) {
					continue
				}
				ip := s[:x] + "." + s[x:x+y] + "." + s[x+y:x+y+z] + "." + s[x+y+z:]
				ipAddress = append(ipAddress, ip)
			}
		}
	}

	return ipAddress
}

func checkIP(str string) bool {
	switch len(str) {
	case 3:
		return str >= "100" && str <= "255"
	case 2:
		return str >= "10" && str <= "99"
	case 1:
		return true
	default:
		return false
	}
}

const SEG_COUNT = 4

var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = []string{}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, segId, segStart int) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == SEG_COUNT {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != SEG_COUNT-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}
