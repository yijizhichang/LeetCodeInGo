package problem

/**
解题思路
1. m位和n位数相乘，结果位数为m+n-1或m+n，所以创建一个m+n位数组
2. 对应位相乘后的结果与进位数据数组对应位置相加，十位数存入进位数组下一位，个位数留在该位 (对10取余数结果为该位数据，进位除以10)
3. 第一位为0的情况，在转为字符串时记得去掉
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	result := make([]int, m+n)
	for j := n - 1; j >= 0; j-- {
		for i := m - 1; i >= 0; i-- {
			tmp := int(num1[i]-'0')*int(num2[j]-'0') + result[i+j+1]
			result[i+j+1] = tmp % 10
			result[i+j] += tmp / 10
		}
	}
	if result[0] == 0 {
		result = result[1:]
	}
	length := len(result)
	str := make([]byte, length)
	for k, val := range result {
		str[k] += byte(val) + '0'
	}
	return string(str)
}
