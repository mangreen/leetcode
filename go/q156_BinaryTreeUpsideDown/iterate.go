package q156_BinaryTreeUpsideDown

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	// 遞迴處理左子樹，直到找到新的根節點 (最底部左節點)
	newRoot := upsideDownBinaryTree(root.Left)

	root.Left.Left = root.Right // 原本的右子樹變成新的左子樹
	root.Left.Right = root // 原本的父節點變成新的右子樹

	// 斷開原本的連結，避免形成環狀結構
	root.Left = nil
	root.Right = nil

	return newRoot
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
