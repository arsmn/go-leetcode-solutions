package main

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var (
		col, row, block [9][9]bool
	)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '1'
			if col[j][n] {
				return false
			}

			if row[i][n] {
				return false
			}

			b := i/3*3 + j/3
			if block[b][n] {
				return false
			}

			col[j][n] = true
			row[i][n] = true
			block[b][n] = true
		}
	}

	return true
}
