package problem

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

 */

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	// 注意越界问题
	if divisor == -1 && dividend == -(math.MaxInt32+1) {
		return math.MaxInt32
	}
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))
	result := div(dividendAbs, divisorAbs)
	// 还原其本来正负
	if (dividend <= 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		return -result
	}
	return result
}

func div(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	sum := 0
	// 只要被除数还比较大就一直循环
	for dividend >= divisor {
		level, originalDivisor, lastDivisor := uint(0), divisor, 0 // X2的次数 / 保存原始的除数 / 保存最后一次够减的数字
		for divisor <= dividend {
			level++
			lastDivisor = divisor
			divisor += divisor
		}
		level-- // 最后一次的X2是超出的，所以-1
		sum += 1 << level // 加上符合的因数
		dividend -= lastDivisor
		divisor = originalDivisor
	}
	return sum
}

