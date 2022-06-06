package lcci

//https://leetcode.cn/problems/binode-lcci/
//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {

	var pre *TreeNode = nil
	dfs(root, &pre)

	return pre
}
func dfs(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return

	}
	dfs(root.Right, pre)
	root.Right = *pre
	if *pre != nil {
		(*pre).Left = nil
	}

	*pre = root
	dfs(root.Left, pre)
	return
}
