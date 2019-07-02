package problem

func minDepth(root *TreeNode) int {
        if root == nil {
                return 0
        }

        if root.Left == nil && root.Right == nil {
                return 1
        }

        if root.Left == nil {
                return minDepth(root.Right) + 1
        }

        if root.Right == nil {
                return minDepth(root.Left) + 1
        }

        return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
        if a > b {
                return b
        }
        return a
}
