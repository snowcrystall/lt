package problems

/**
 * Definition for a Node.
 */

type Node struct {
	Val      int
	Children []*Node
}

//429 n-ary-tree-level-order-traversal
//给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}
		res = append(res, level)
	}
	return res

}
