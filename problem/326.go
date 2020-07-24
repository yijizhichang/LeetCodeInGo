package problem

/*
3 的 幂
*/

func isPowerOfThree(n int) bool {
	//设计思路，比较n和3的幂的大小，若n < 1 (3**0) 则返回false，n ==1 返回true
	// n > 3 ,进入下次循环，比较 n < 3 (3**1)
	aa := 1
	for {
		if n < aa {
			return false
		}
		if n == aa {
			return true
		}
		aa *= 3
	}

}
