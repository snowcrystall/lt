package problems

//Definition for a QuadTree node.
type QNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QNode
	TopRight    *QNode
	BottomLeft  *QNode
	BottomRight *QNode
}

func construct(grid [][]int) *QNode {
	var dfs func([][]int, int, int) *QNode
	dfs = func(rows [][]int, c0, c1 int) *QNode {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] { // 不是叶节点
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &QNode{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		// 是叶节点
		return &QNode{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}
