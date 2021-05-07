package array_list

/*
如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A是单调数组时返回 true，否则返回 false。

*/

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	upSingleCount := 0
	downSingleCount := 0
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			upSingleCount++
		} else if A[i] < A[i-1] {
			downSingleCount++
		} else {
			upSingleCount++
			downSingleCount++
		}
	}
	return upSingleCount == len(A)-1 || downSingleCount == len(A)-1
}
