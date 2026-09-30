func isValidSudoku(board [][]byte) bool {
    squaresIdxs := [9][2]int{
        [2]int{0,0},
        [2]int{0,3},
        [2]int{0,6},
        [2]int{3,0},
        [2]int{3,3},
        [2]int{3,6},
        [2]int{6,0},
        [2]int{6,3},
        [2]int{6,6},
    }
    for i, row := range board {
        if !isValidRow(row) || !isValidRow(columnToRow(board, i)) || !isValidRow(squareToRow(board, squaresIdxs[i])){
            return false
        }   
    }
    return true
}

func columnToRow(board [][]byte, c int) []byte {
    row := []byte{}
    for r, _ := range board{
        row = append(row, board[r][c])
    }
    return row
}

func squareToRow(board [][]byte, idxs [2]int) []byte{
    row := []byte{}
    rIdx := idxs[0]
    cIdx := idxs[1]
    for r := rIdx; r < (rIdx + 3); r ++ {
        for c := cIdx; c < (cIdx + 3); c++ {
            row = append(row, board[r][c])
        }
    }
    return row
}

func isValidRow(row []byte) bool {
    seen := make(map[rune]struct{})
    for _, b := range row{
        char := rune(b)
        if _, ok := seen[char]; char != '.' && ok{
            return false
        }
        seen[char] = struct{}{}
    } 
    return true
}
