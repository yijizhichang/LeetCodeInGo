package search

func minDays(bloomDay []int, m, k int) int {
	if k*m > len(bloomDay) {
		return -1
	}
	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}
	left, right := 0, maxDay
	mid := 0
	for left <= right {
		mid = right - (right-left)/2
		if canMake(bloomDay, mid, m, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canMake(bloomDay []int, days, m, k int) bool {
	bouquets := 0
	canUsedFlowers := 0
	for _, d := range bloomDay {
		if d > days {
			canUsedFlowers = 0
		} else {
			canUsedFlowers++
			if canUsedFlowers == k {
				bouquets++
				canUsedFlowers = 0
			}
		}
	}
	return bouquets >= m
}
