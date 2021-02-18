/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return rserialize(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if data == "" {
		return nil
	}

	vals := strings.Split(data, ",")
	result := rdeserialize(&vals)

	return result

}

func rdeserialize(vals *[]string) *TreeNode {

	if (*vals)[0] == "nil" {
		*vals = (*vals)[1:]
		return nil
	}

	root := new(TreeNode)
	root.Val, _ = strconv.Atoi((*vals)[0])
	*vals = (*vals)[1:]
	root.Left = rdeserialize(vals)
	root.Right = rdeserialize(vals)
	return root
}

func rserialize(root *TreeNode) string {

	if root == nil {
		return "nil"
	}
	result := strconv.Itoa(root.Val) + "," + rserialize(root.Left) + "," + rserialize(root.Right)
	return result
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

