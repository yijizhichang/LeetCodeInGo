package xor

/*
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。

对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。

并返回一个包含给定查询 queries 所有结果的数组。



示例 1：

输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8]
解释：
数组中元素的二进制表示形式是：
1 = 0001
3 = 0011
4 = 0100
8 = 1000
查询的 XOR 值为：
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8

*/

func xorQueries(arr []int, queries [][]int) []int {
	ret := make([]int, 0)
	for _, query := range queries {
		this := arr[query[0]]
		for i := query[0] + 1; i <= query[1]; i++ {
			this ^= arr[i]
		}
		ret = append(ret, this)
	}
	return ret
}

/*

当 left=0 时，Q(left,right)=xors[right+1]

当 left>0 时，Q(left,right) 的计算如下：
 Q(left,right)
=arr[left]⊕…⊕arr[right]
=(arr[0]⊕…⊕arr[left−1])⊕(arr[0]⊕…⊕arr[left−1])⊕(arr[left]⊕…⊕arr[right])
=(arr[0]⊕…⊕arr[left−1])⊕(arr[0]⊕…⊕arr[right])
=xors[left]⊕xors[right+1]

*/
// 前缀和
func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = xors[q[0]] ^ xors[q[1]+1]
	}
	return ans
}
