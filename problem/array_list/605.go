package array_list
/*
605. 种花问题
假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。

示例 1:

输入: flowerbed = [1,0,0,0,1], n = 1
输出: True
示例 2:

输入: flowerbed = [1,0,0,0,1], n = 2
输出: False
注意:

数组内已种好的花不会违反种植规则。
输入的数组长度范围为 [1, 20000]。
n 是非负整数，且不会超过输入数组的大小。
 */

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if n == 0 {
		return true
	}
	if length == 1 {
		return flowerbed[0] == 0 && n == 1
	}
	for idx, isFlower := range flowerbed {
		if isFlower == 0 {
			canFlower := false
			// 判断前后是否可以种植
			if idx == 0 {
				if length > 1 && flowerbed[idx+1] != 1 {
					canFlower = true
				}
			} else if idx == length - 1 {
				if length > 1 && flowerbed[idx-1] != 1 {
					canFlower = true
				}
			} else {
				if flowerbed[idx-1] != 1 && flowerbed[idx+1] != 1 {
					canFlower = true
				}
			}
			if canFlower {
				flowerbed[idx] = 1
				n--
			}
			if n == 0 {
				break
			}
		}
	}
	return n == 0
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for idx := 0 ; idx < len(flowerbed) ; idx += 2{
		if flowerbed[idx] == 0 {
			// 如果是最后一格或者下一格为空,就可以种上
			if idx == len(flowerbed) - 1 || flowerbed[idx+1] == 0 {
				n --
			} else {
				idx ++
			}
		}
	}
	return n<=0
}
