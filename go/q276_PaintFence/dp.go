package q276_PaintFence

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	
	if n == 1 {
		return k
	}

	// 初始狀態: 第二欄
	same := k
	diff := k * (k - 1)
	all := same + diff

	// 動態規劃: 從第三欄開始計算
	for i:=3; i<=n; i++ {
		same = diff
		diff = all * (k - 1)

		all = same + diff
	}
	
	return all
}
/*
same: 
..., (i-2), (i-1), i, ... 
第i-1欄跟第i欄顏色相同, 第i-1欄顏色要和第i-2欄不同, 
所以有第i-1欄的diff種選擇

diff:
..., (i-1), i, ... 
第i-1欄有k種, 第i欄顏色要和第i-1欄不同, 所以有k-1種選擇
所以有all * (k-1)種選擇

ex.
n=1 k=2 ⭢ 2
R or B

n=2 k=2 ⭢ 4
RR, RB, BB, BR

n=3 k=2 ⭢ 6
RRR, RRB, RBB, BBR, BBR, BRR

n=4 k=2 ⭢ 10
RRRR, RRRB, RRBB, RBBR, BBBB, BBBR, BBRR, BRRR, RBRB, BRBR
*/