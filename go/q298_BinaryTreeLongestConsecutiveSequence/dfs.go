package q298_BinaryTreeLongestConsecutiveSequence

func longestConsecutive(root *TreeNode) int {
	ans := 0
	DFS(root, 0, 0, &ans)
	return ans
}

func DFS(root *TreeNode, parentVal int, curLen int, maxLen *int) {
	if root == nil {
		return
	}

	if root.Val == parentVal+1 {
		curLen++
	} else {
		curLen = 1
	}

	*maxLen = max(*maxLen, curLen)

	DFS(root.Left, root.Val, curLen, maxLen)
	DFS(root.Right, root.Val, curLen, maxLen)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
