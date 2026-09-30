func hasDuplicate(nums []int) bool {
    numsSeen := make(map[int]struct{})
    for _, n := range nums {
        _, exists := numsSeen[n]
        if exists {
            return true
        }
        numsSeen[n] = struct{}{}
    }
    return false
}
