package sort

func QuickSort(tar []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := Partition(tar, low, high)
	QuickSort(tar, low, pivotIndex-1)
	QuickSort(tar, pivotIndex+1, high)
}

// 双边循环
func Partition(tar []int, low, high int) int {
	left, right := low, high
	pivot := tar[left]
	// 终止条件是左右指向同一个值
	for left != right {
		// 一定先找右边的
		// 右边开始，找到第一个小于哨兵值的点
		for left < right && pivot <= tar[right] {
			right--
		}
		// 从左边开始，找到第一个大于哨兵值的点
		for left < right && pivot >= tar[left] {
			left++
		}
		if left < right {
			// 交换
			tar[left], tar[right] = tar[right], tar[left]
		}
	}
	// 把哨兵放到最终满足左边都小于他，右边都大于他的点
	tar[low], tar[left] = tar[left], tar[low]
	return left
}
