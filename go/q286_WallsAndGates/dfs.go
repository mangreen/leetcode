package q286_WallsAndGates

import ()

func wallsAndGates(rooms [][]int) {
    for i, row := range rooms {
        for j, x := range row {
            // 從門 (0) 開始 DFS
            if x == 0 {
                DFS(rooms, i, j, 0)
            }
        }   
    }
}


func DFS(rs [][]int, i, j, dist int) {
    if i < 0 || i >= len(rs) || j < 0 || j >= len(rs[0]) {
        return
    }

    // 遇到牆或門，停止
    if rs[i][j] < dist {
        return
    }

    // 更新距離
    rs[i][j] = dist
    
    // 四個方向繼續 DFS
    DFS(rs, i+1, j, dist+1)
    DFS(rs, i-1, j, dist+1)
    DFS(rs, i, j+1, dist+1)
    DFS(rs, i, j-1, dist+1)
}