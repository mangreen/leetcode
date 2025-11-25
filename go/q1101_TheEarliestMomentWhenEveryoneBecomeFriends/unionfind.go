package q1101_TheEarliestMomentWhenEveryoneBecomeFriends

import (
	"slices"
	"sort"
)

func earliestAcq(logs [][]int, n int) int {
	// 轉換日誌為自定義結構體便於排序
	cs := make([]Connection, len(logs))
	for i, log := range logs {
		cs[i] = Connection{
			timestamp: log[0],
			x:         log[1],
			y:         log[2],
		}
	}

	sort.Slice(cs, func(i, j int) bool {
		return cs[i].timestamp < cs[j].timestamp
	})

	// Go 1.23+ 新增 lices.Repeat
	roots := slices.Repeat([]int{-1}, n)
	cnt := slices.Repeat([]int{1}, n) // 每個朋友圈一開始只有自己

	for _, c := range cs {
		x := find(roots, c.x)
		y := find(roots, c.y)

		if x != y {
			// 合併較小的集合到較大的集合中
			if cnt[x] < cnt[y] {
				x, y = y, x
			}

			roots[y] = x
			cnt[x] += cnt[y]
		}

		// 檢查最大集合是否包含所有人
		if cnt[x] == n {
			return c.timestamp
		}
	}

	return -1
}

// 定義連接事件的結構體
type Connection struct {
	timestamp int
	x, y      int
}

func find(roots []int, cur int) int {
	for roots[cur] != -1 {
		cur = find(roots, roots[cur])
	}

	return cur
}
