func searchMatrix(matrix [][]int, target int) bool {
    firstCol := getFirstCol(matrix)
    rowIdx, found := binarySearch(firstCol, target)
    if rowIdx == -1 {
        return false
    }
    _, found = binarySearch(matrix[rowIdx], target)
    return found


}

/*

[1 , 3, 5, 7]
[10,11,16,20]
[23,30,34,60]

*/

func getFirstCol(matrix [][]int) []int {
    col := []int{}
    for _, row := range matrix{
        col = append(col, row[0])
    }

    return col
}

func binarySearch(row []int, target int) (int, bool) {
    l := 0
    r := len(row) - 1
    var m int
    for l <= r {
        m = int((l + r) / 2)
        if row[m] == target {
            return m, true
        } else if row[m] < target {
            l = m+1
        } else {
            r = m-1
        }
    }

    if target < row[m] {
        if m == 0{
            return -1, false
        } else{
            return m-1, false
        }
    } else {
        return m, false
    }

}
