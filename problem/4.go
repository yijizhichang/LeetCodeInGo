package problem

func hmax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hmin(x,y int) int {
	if x>y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	l:=0
	r:=len(height)-1
	max := 0
	for l<r {
		max = hmax(max, hmin(height[r],height[l])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
