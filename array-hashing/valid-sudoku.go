package arrayHashing

import "fmt"

// Time Complexity O(n^2)
// Space Complexity O(n)
func IsValidSudoku(board [][]byte) bool {
	hashMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := i
			column := j

			current_val := string(board[i][j])

			if current_val == "." {
				continue
			}

			foundInRow := fmt.Sprintf("%v found in row %v", current_val, row)
			foundInColumn := fmt.Sprintf("%v found in column %v", current_val, column)
			foundInGrid := fmt.Sprintf("%v found in grid %v-%v", current_val, i/3, j/3)

			_, okRow := hashMap[foundInRow]
			_, okCol := hashMap[foundInColumn]
			_, okGrid := hashMap[foundInGrid]

			if okRow || okCol || okGrid {
				return false
			} else {
				println(foundInRow)
				println(foundInColumn)
				println(foundInGrid)

				hashMap[foundInRow] = true
				hashMap[foundInColumn] = true
				hashMap[foundInGrid] = true
			}
		}
	}

	return true
}
