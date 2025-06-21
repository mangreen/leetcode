package q323_NumberOfConnectedComponentsInAnUndirectedGraph

func countComponents3(n int, edges [][]int) int {
    ans := n

    // 初始 n 個 node 有 n 個 union
    roots := make([]int, n)
    for i:=0; i<n; i++ {
        roots[i] = -1
    }

    for _, e := range edges {
        x := find(roots, e[0])
        y := find(roots, e[1])

        // 每找(加) 1 條邊/加入 1 個 node = 少 1 個 node(union)
        if x != y {
            ans--
            roots[x] = y
        }
    }

    return ans
}

func find(roots []int, cur int) int {
    for roots[cur] != -1 {
        cur = find(roots, roots[cur])
    }

    return cur
}
/*
Time: O(N)
Space: O(N)

countComponents3(3, [][]int{{0, 1}, {1, 2}, {0, 2}})

ans=n=3
e[0 1] x=0 y=1 roots[1 -1 -1] ans=2
e[1 2] x=1 y=2 roots[1 2 -1] ans=1
e[0 2] x=2 y=2 roots[1 2 -1] ans=1
ANS: 1
*/