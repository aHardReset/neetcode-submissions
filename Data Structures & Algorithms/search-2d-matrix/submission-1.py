class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        rows = len(matrix)
        cols = len(matrix[0])
        left = 0
        right = (rows*cols)-1
        while left <= right:
            middle = (left + right) // 2
            i,j = divmod(middle, cols)
            value = matrix[i][j]
            if value == target:
                return True
            elif value < target:
                left = middle + 1
            else:
                right = middle - 1
        return False

        