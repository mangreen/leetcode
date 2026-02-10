package q250_CountUnivalueSubtrees


func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	cnt := 0
	if isUnival(root, root.Val) {
		cnt++
	}
	
	return cnt + countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
}

func isUnival(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	
	if root.Val != val {
		return false
	}
	
	return isUnival(root.Left, val) && isUnival(root.Right, val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
