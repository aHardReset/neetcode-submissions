class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # print(self.areValidRows(board))
        # print(self.areValidCols(board))
        # print(self.areValidSquares(board))
        return self.areValidRows(board) and self.areValidCols(board) and self.areValidSquares(board)

    def areValidSquares(self, board: List[List[str]]):
        for i in range(0, 9, 3):
            for j in range(0, 9, 3):
                nums = []
                for _i in range(i, i+3):
                    for _j in range(j, j+3):
                        if board[_i][_j].isdigit():
                            nums.append(board[_i][_j])
                if not self.areDistinctNumbers(nums):
                    return False
        return True

    def areValidCols(self, board: List[List[str]]) -> bool:
        copied = []
        for i in range(len(board)):
            col = []
            for j in range(len(board[0])):
                col.append(board[j][i])
            copied.append(col)
        return self.areValidRows(copied)


    def areValidRows(self, board: List[List[str]]) -> bool:
        for row in board:
            nums = []
            for num in row:
                if num.isdigit():
                    nums.append(num)
            if not self.areDistinctNumbers(nums):
                return False
        return True

    def areDistinctNumbers(self, nums: List[int]) -> bool:
        seen = []
        for num in nums:
            if num in seen:
                return False
            seen.append(num)
        return True
        