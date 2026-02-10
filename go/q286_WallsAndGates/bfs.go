package q286_WallsAndGates

import ()

func wallsAndGatesBFS(rooms [][]int) {
    q := []pair{}

    for i, row := range rooms {
        for j, x := range row {
            // 從門 (0) 開始 BFS
            if x == 0 {
                q = append(q, pair{x: i, y: j})
            }
        }   
    }

    dirs := []pair{{1,0}, {-1,0}, {0,1}, {0,-1}}

    for len(q) > 0 {
        for _, cur := range q {
            q = q[1:]

            for _, dir := range dirs {
                nx, ny := cur.x+dir.x, cur.y+dir.y

                // 超出邊界或不是2147483647(已經更新過/牆/門)，跳過
                if nx < 0 || nx >= len(rooms) || ny < 0 || ny >= len(rooms[0]) || rooms[nx][ny] != 2147483647 {
                    continue
                }

                rooms[nx][ny] = rooms[cur.x][cur.y] + 1
                q = append(q, pair{x: nx, y: ny})
            }
        }
    }
}

type pair struct {
    x, y int
}