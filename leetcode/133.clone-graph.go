/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	visited := make(map[*Node]*Node)
	return clone(node, visited)

	// recursive dfs
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, found := visited[node]; found {
		return v
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode

}

// @lc code=end

