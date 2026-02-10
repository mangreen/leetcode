package q270_ClosestBinarySearchTreeValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	DFS(root, target, &closest)

	return closest
}

func DFS(n *TreeNode, target float64, closest *int) {
	if n == nil {
		return
	}

	if abs(float64(n.Val)-target) < abs(float64(*closest)-target) {
		*closest = n.Val
	}

	if target < float64(n.Val) {
		DFS(n.Left, target, closest)
	}

	if target > float64(n.Val) {
		DFS(n.Right, target, closest)
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

/*
ex1.
	   4
	  / \
	 2   5
	/ \
   1   3

target=3.714286

root=4=closest ⭢ 4-3.714286=0.285714 < 4-2=2 ⭢ 更新closest=4
target=3.714286 < 4 ⭢ 左子树

root=2 ⭢ 3.714286 - 2=1.714286 > 4-2=2 ⭢ 不更新closest
target=3.714286 > 2 ⭢ 右子树

root=3 ⭢ 3.714286 - 3=0.714286 > 4-3=0.285714 ⭢ 不更新closest
target=3.714286 > 3 ⭢ 右子树为空，结束

return closest=4

ex2.
   1
target=1.0

root=1=closest ⭢ 1-1.0=0 < 1-1=0 ⭢ 不更新closest
target=1.0 == 1 ⭢ 右子树为空，结束
*/
