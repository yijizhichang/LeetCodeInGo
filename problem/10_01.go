package problem

//
//func merge(A []int, m int, B []int, n int){
//	for kb, vb := range B {
//		if m > 0 {
//			for ka := kb + m - 1; ka >= 0 ; ka -- {
//				if A[ka] <= vb {
//					// B 元素大于A元素 将B元素放在该A元素后面
//					A[ka+1] = vb
//					break
//				} else {
//					// B 元素小于 A元素 将A元素向后移一位 B插入当前A位置
//					A[ka+1] = A[ka]
//					A[ka] = vb
//				}
//			}
//		} else {
//			A[kb] = vb
//		}
//	}
//}

func merge(A []int, m int, B []int, n int) {
	idx := len(A) - 1
	a, b := m-1, n-1
	for a >= 0 && b >= 0 {
		if A[a] > B[b] {
			A[idx] = A[a]
			a--
		} else {
			A[idx] = B[b]
			b--
		}
		idx--
	}
	for a >= 0 {
		A[idx] = A[a]
		a--
		idx--
	}
	for b >= 0 {
		A[idx] = B[b]
		b--
		idx--
	}
}
