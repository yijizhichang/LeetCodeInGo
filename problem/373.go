package problem

import (
	"fmt"
	"sort"
)

/*

给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。

找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。

示例 1:

	输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
	输出: [1,2],[1,4],[1,6]
	解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:

	输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
	输出: [1,1],[1,1]
	解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
示例 3:

	输入: nums1 = [1,2], nums2 = [3], k = 3
	输出: [1,3],[2,3]
	解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]


*/

/*
	解题：
	解法法一 ：kSmallestPairs
	最先想到用最直接的方式解决，把所有的情况放到二位数据里
	然后对二维数组按照 和排序
	然后取出 k对或者 len(nums1) * len(nums2)对数组

	最后用了 插入排序  超时不通过

	然后用了  map[int][][]int{}  这个结构  map 的key 存储 u+v 的值

	将key 的值 放进 []int  然后排序

	循环依次拿出 map 里的值  这个通过了

	解法二：kSmallestPairsFast

	解法二用的是堆排序做的，时间复杂度 nlogn 速度很快

	大顶堆的特点

	root节点为当前堆中最大的值
	是二叉树，并不是平衡二叉树
	父节点一定比子节点大

	当前节点为index
	parent节点 为 (index - 1)/2
	leftChild 为	  index*2 + 1
	rightChjld 为  index*2 + 2

	每次添加或删除节点

	重置节点 保证 root 节点为最大值


*/

/*
func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}

	nums1 = []int{-476570184, -423568801, -385585840, -375390924, -364630569, -359795128, -281872968, -126410430, -75677925, -54214495, -49178055, -32637211, -32198215, 3413177, 19045759, 62248526, 67551536, 113606647, 155411580, 164755463, 164781059, 203133270, 277305105, 284913246, 285973110, 296436629, 325431544, 357294459, 378678394, 399786157}
	nums2 = []int{-408663357, -404578641, -376531700, -311642519, -294905976, -232001207, -183530032, -141524508, -115652480, -70696522, -63386299, -54656543, -32316999, 29714175, 33993996, 45020708, 62165363, 84210823, 93905151, 102177224, 209285622, 288668099, 328300713, 338684779, 342861859, 384940859, 408019604, 410097843, 458721542, 475395296}
	//1000

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}


	//[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

	a := kSmallestPairsFast(nums1, nums2, 1000)

	fmt.Println(a)

}
*/

/*
	方法一
*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var arr [][]int
	var ar []int
	if k <= 0 {
		return arr
	}
	lens1 := len(nums1)
	lens2 := len(nums2)
	if lens1 == 0 || lens2 == 0 {
		return arr
	}
	if k > lens1*lens2 {
		k = lens1 * lens2
	}
	a := map[int][][]int{}

	for _, v := range nums1 {
		for _, val := range nums2 {
			ar = append(ar, v+val)
			a[v+val] = append(a[v+val], []int{v, val})
		}
	}
	sort.Ints(ar)

	ar = ar[0:k]
	for _, v := range ar {

		val := a[v][0]

		arr = append(arr, val)

		a[v] = a[v][1:]
		fmt.Println(a[v])
	}
	return arr
}

/*
	方法二
*/
func kSmallestPairsFast(nums1 []int, nums2 []int, k int) [][]int {
	n := len(nums1)
	m := len(nums2)
	if m == 0 || n == 0 || k == 0 {
		return nil
	}
	heap := MaxHeap{
		arr: make([]*Node, k),
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if heap.size < k {
				heap.Add(&Node{
					a:   i,
					b:   j,
					Val: nums1[i] + nums2[j],
				})
			} else {
				if heap.Top().Val > nums1[i]+nums2[j] {
					heap.Pop()
					heap.Add(&Node{
						a:   i,
						b:   j,
						Val: nums1[i] + nums2[j],
					})
				} else {
					break
				}
			}
		}
	}
	ret := make([][]int, heap.Size())
	i := heap.Size() - 1
	for !heap.IsEmpty() {
		node := heap.Pop()
		ret[i] = []int{nums1[node.a], nums2[node.b]}
		i--
	}
	return ret
}

/*
	节点结构体 存储  u,v u+v
*/
type Node struct {
	a   int // nums1[n]
	b   int // nums2[n]
	Val int // nums1[n] + nums2[n]
}

/*
	大顶堆结构体 MaxHeap[0] 为当前堆中最大值
*/

type MaxHeap struct {
	arr  []*Node
	size int // 最大为k 或者 为 len(nums1) * len(nums2
}

/*
	为大顶堆添加 节点 并调整堆的 root 节点为最大值
*/

func (h *MaxHeap) Add(n *Node) {
	h.arr[h.size] = n
	pos := h.size
	h.size++
	for pos > 0 {
		parent := h.Parent(pos)
		if h.arr[parent].Val > n.Val {
			break
		}
		h.Swap(parent, pos)
		pos = parent
	}
}

/*
	踢出堆的最大值
*/
func (h *MaxHeap) Pop() *Node {
	if h.IsEmpty() {
		return nil
	}

	ret := h.arr[0]
	h.size--
	h.arr[0] = h.arr[h.size]
	pos := 0
	/*
		踢出 堆中最大的值之后 从root节点开始 依次比对左右子节点
		子节点中最大的值 大于 父节点  与父节点调换
		重置堆保证root节点为堆中最大值
	*/
	for pos < h.size {
		left, right := h.Child(pos)
		if (left != -1 && h.arr[left].Val > h.arr[pos].Val) || (right != -1 && h.arr[right].Val > h.arr[pos].Val) {
			if right != -1 && h.arr[left].Val < h.arr[right].Val {
				h.Swap(right, pos)
				pos = right
			} else {
				h.Swap(left, pos)
				pos = left
			}
		} else {
			break
		}
	}
	return ret
}

func (h *MaxHeap) Swap(a, b int) {
	tmp := h.arr[a]
	h.arr[a] = h.arr[b]
	h.arr[b] = tmp
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) Size() int {
	return h.size
}

func (h *MaxHeap) Top() *Node {
	if h.IsEmpty() {
		return nil
	}
	return h.arr[0]
}

func (h *MaxHeap) Parent(pos int) int {
	return (pos - 1) / 2
}

func (h *MaxHeap) Child(pos int) (int, int) {
	left := pos*2 + 1
	right := pos*2 + 2
	if left >= h.size {
		left = -1
	}
	if right >= h.size {
		right = -1
	}
	return left, right
}
