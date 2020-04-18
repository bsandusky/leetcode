package minPathSum

func minPathSum(grid [][]int) int {

	col, row := len(grid), len(grid[0])
	for i := 1; i < col; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for j := 1; j < row; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[col-1][row-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
