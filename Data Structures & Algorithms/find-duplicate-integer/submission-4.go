func findDuplicate(nums []int) int {
    // O(n), O(n)
    seen := map[int]struct{}{} // map[int]struct{} <- def {} <- initial value

    for _, num := range nums {
        if _, ok := seen[num]; ok {
            return num
        }
        seen[num] = struct{}{}
    }
    return 0
}
