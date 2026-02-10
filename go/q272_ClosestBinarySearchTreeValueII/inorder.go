package q270_ClosestBinarySearchTreeValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	ans := []int{}
	
	inorder(root, target, k, &ans)

	return ans
}

func inorder(n *TreeNode, target float64, k int, ans *[]int) {
	if n == nil {
		return
	}

	// 中序遍歷，左⭢根⭢右
	inorder(n.Left, target, k, ans)

	if len(*ans) < k {
		*ans = append(*ans, n.Val)
	} else {
		// 已經有k個值了，檢查是否需要更新
		if abs(float64(n.Val)-target) < abs(float64((*ans)[0])-target) {
			*ans = (*ans)[1:] // 移除最遠的值
			*ans = append(*ans, n.Val)
		} else {
			// 提前剪枝，因為中序遍歷是排序好的，後面的值只會更大
			return
		}
	}

	inorder(n.Right, target, k, ans)
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
target=3.714286, k=2

root=4 ⭢ 左子樹	

root=2 ⭢ 左子樹

root=1 ⭢ 左子樹为空，返回
ans=[] < k=2 ⭢ 添加1 ans=[1]
root=1 ⭢ 右子树为空，返回

root=2 ⭢ ans=[1] < k=2 ⭢ 添加2 ans=[1,2]
root=2 ⭢ 右子树

root=3 ⭢ 左子树为空，返回
ans=[1,2] == k=2 ⭢ 检查是否需要更新
|3-3.714286|=0.714286 < |1-3.714286|=2.714286 ⭢ 更新 ans=[2,3]
root=3 ⭢ 右子树为空，返回

root=4 ⭢ ans=[2,3] == k=2 ⭢ 检查是否需要更新
|4-3.714286|=0.285714 < |2-3.714286|=1.714286 ⭢ 更新 ans=[3,4]
root=4 ⭢ 右子树

root=5 ⭢ 左子树为空，返回
ans=[3,4] == k=2 ⭢ 检查是否需要更新
|5-3.714286|=1.285714 < |3-3.714286|=0.714286 ⭢ 不更新 ans=[3,4] 
提前剪枝，结束

return ans=[3,4]

ex2.
	1
	 \
	  2
target=1.0, k=1

root=1 ⭢ 左子树为空，返回
ans=[] < k=1 ⭢ 添加1 ans=[1]
root=1 ⭢ 右子树

root=2 ⭢ 左子树为空，返回
ans=[1] == k=1 ⭢ 检查是否需要更新
|2-1.0|=1.0 < |1-1.0|=0.0 ⭢ 不更新 ans=[1]
提前剪枝，结束

return ans=[1]
*/
