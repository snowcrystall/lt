package problems

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res

}

// 100.  same-tree
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)

}

//101. 对称二叉树 给你一个二叉树的根节点 root ， 检查它是否轴对称
func isSymmetric(root *TreeNode) bool {
	return dfscheck(root.Left, root.Right)
}

func dfscheck(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && dfscheck(left.Left, right.Right) && dfscheck(left.Right, right.Left)
}

//104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	root := &TreeNode{}
	if len(nums) == 0 {
		return nil
	}
	root.Val = nums[len(nums)/2]
	root.Left = sortedArrayToBST(nums[0 : len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1 : len(nums)])
	return root
}

//110. 平衡二叉树 给定一个二叉树，判断它是否是高度平衡的二叉树。
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

//111. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
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

	if a < b {
		return a
	}
	return b
}
