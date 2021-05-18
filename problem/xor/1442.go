package xor

/*
给你一个整数数组 arr 。

现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。

a 和 b 定义如下：

a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
注意：^ 表示 按位异或 操作。

请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。

 */

/*
   ！！！！
	异或很重要的一点就是前缀和 + 两个相同值异或必然为0 + 任意值和0异或必然为该值
	通过前缀和去推算公示
 */

// 看图1442 可知 问题就是找满足si = s(k+1) 时，i,j,k三元组总数

// 两重循环
func countTriplets(arr []int) int {
	n := len(arr)
	ans := 0
	// 多取一位，第一位=0没有影响
	s := make([]int, n+1)
	for idx, this := range arr {
		s[idx+1] = s[idx]^this
	}
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if s[i] == s[k+1] {
				ans += k - i
			}
		}
	}
	return ans
}

// 一重循环
/*
对于下标k， 若下标 i=i1,i2,...,im时均满足 Si = S(k+1),根据二重循环，这些二元组(i1,k),(i2,k),...,(im,k)对答案的贡献之和为
(k-i1) + (k-i2) + ... + (k-im) = m*k - (i1+i2+...+im)

也就是说，当遍历下标k时，我们需要知道所有满足Si=S(k+1)的：
下标i的出现次数m
下标i之和

 */

func countTriplets(arr []int) (ans int) {
	n := len(arr)
	s := make([]int, n+1)
	for i, v := range arr {
		s[i+1] = s[i] ^ v
	}
	cnt := map[int]int{}
	total := map[int]int{}
	for k := 0; k < n; k++ {
		if m, has := cnt[s[k+1]]; has {
			ans += m*k - total[s[k+1]]
		}
		cnt[s[k]]++
		total[s[k]] += k
	}
	return
}

