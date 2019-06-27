package problem

/**
 最小高度树
 解题思路：最小高度树的特点是根节点位于某条最长路径的中点，中点可能是一个，也有可能是两个
 通过一层一层地将边缘的叶子节点去掉之后，最后保存下来的一个或者两个点就是结果。叶子节点的特点是只有一条边。
 */

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	nodes := make([][]int, n)
	for i:=0; i < n; i ++ {
		nodes[i] = make([]int, 0, 3)
	}
	count := make([]int, n)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
		count[edge[0]] ++
		count[edge[1]] ++
	}
	leaves := make([]int, 0)
	for i := 0; i <n ;i++ {
		if count[i] == 1 {
			leaves = append(leaves, i)
		}
	}
	var leafNode, leaf int
	for n > 2 {
		newLeaves := make([]int, 0, len(leaves))
		for _, leaf = range leaves {
			n --
			for _, leafNode = range nodes[leaf] {
				count[leafNode] --
				if count[leafNode] == 1 {
					newLeaves = append(newLeaves, leafNode)
				}
			}
		}
		leaves = newLeaves
	}
	return leaves
}
