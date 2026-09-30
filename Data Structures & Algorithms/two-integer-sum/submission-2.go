func twoSum(nums []int, target int) (z []int) {
    numMap := map[int]int{} // {4:0, }
    for idx2, n := range nums {
        // 1, 4
        if idx1, ok := numMap[n] ; ok{
            return []int{idx1, idx2}
        }
        numMap[target-n] = idx2
    }
    return z
}
