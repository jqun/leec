package l934

// 扩展边缘层数BFS
func shortestBridge(grid [][]int) (step int) {
	n := len(grid)
	for i, row := range grid {
		for j, val := range row {
			if val != 1 {
				continue // 未曾找到岛
			}
			var island []pair
			grid[i][j] = -1     // 标记
			q := []pair{{i, j}} // 以此为起点搜索全部的1--BFS
			for len(q) > 0 {    // 找到此1连通的所有1，即为找打岛屿1
				p := q[0] // 取出剩余节点的第一个节点
				q = q[1:] // 其他所有节点
				island = append(island, p)
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y // 相邻节点
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						q = append(q, pair{x, y})
					}
				}
			}
			// 找到完整岛屿
			q = island
			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, dir := range dirs {
						x, y := p.x+dir.x, p.y+dir.y
						if x >= 0 && x < n && y >= 0 && y < n {
							if grid[x][y] == 1 { // 找到相邻点
								return
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1           // 避免重复搜索
								q = append(q, pair{x, y}) // BFS
							}
						}
					}
				}
				step++ // 一圈搜索完毕
			}
		}
	}
	return
}

// 方向定义
type pair struct {
	x int
	y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
