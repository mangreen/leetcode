package q261_GraphValidTree

import "slices"

func validTree2(n int, edges [][]int) bool {
	// Go 1.23+ 新增 lices.Repeat
	roots := slices.Repeat([]int{-1}, n)

	for _, e := range edges {
		x := find(roots, e[0])
		y := find(roots, e[1])

		// 已在同一集合 -> cycle
		if x == y {
			return false
		}

		roots[x] = y
	}

	// 樹的邊數量必定比節點數量少1
	return len(edges) == n-1
}

func find(roots []int, cur int) int {
	// 當 roots[cur] == -1 時, 表示 cur 為該集合的 root, 或還沒有指向其他節點
	for roots[cur] != -1 {
		cur = find(roots, roots[cur])
	}

	return cur
}
/*
ex1.
n=4
edges=[[0,1],[0,2],[0,3]]
1 ― 0 ― 3
    | 
    2
roots=[-1,-1,-1,-1]

e:[0,1] ⭢ x=find(0)=0, y=find(1)=1 ⭢ roots=[1,-1,-1,-1], 把 0 指向 1
e:[0,2] ⭢ x=find(0)=1, y=find(2)=2 ⭢ roots=[1,2,-1,-1], 把 1 指向 2
e:[0,3] ⭢ x=find(0)=2, y=find(3)=3 ⭢ roots=[1,2,3,-1], 把 2 指向 3
return len(edges)=3 == n-1=3 ⭢ true

ex2.
n=3
edges=[[0,1],[0,2],[0,2]]
  0 
 / \ 
1 ― 2
roots=[-1,-1,-1]

e:[0,1] ⭢ x=find(0)=0, y=find(1)=1 ⭢ roots=[1,-1,-1], 把 0 指向 1
e:[0,2] ⭢ x=find(0)=1, y=find(2)=2 ⭢ roots=[1,2,-1], 把 1 指向 2
e:[0,2] ⭢ x=find(0)=2, y=find(2)=2 ⭢ 0 跟 2 在同一集合, 有迴圈 -> return false
*/
