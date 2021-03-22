package matrix


func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i:=0;i<n;i++ {
		ret[i] = make([]int, n)
	}
	i:=1
	leftBound, upBound := 0, 0
	rightBound, DownBound := n-1, n-1
	for leftBound <= rightBound && upBound <= DownBound {
		for x:=leftBound;x<=rightBound;x++ {
			ret[upBound][x] = i
			i++
		}
		upBound++
		if upBound > DownBound {
			break
		}
		for y:=upBound;y<=DownBound;y++ {
			ret[y][rightBound] = i
			i++
		}
		rightBound--
		if rightBound < leftBound{
			break
		}
		for x:=rightBound;x>=leftBound;x-- {
			ret[DownBound][x] = i
			i++
		}
		DownBound--
		if DownBound < upBound {
			break
		}
		for y:=DownBound;y>=upBound;y-- {
			ret[y][leftBound] = i
			i++
		}
		leftBound++
	}
	return ret
}
