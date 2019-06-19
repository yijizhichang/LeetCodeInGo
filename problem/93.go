package problem

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
