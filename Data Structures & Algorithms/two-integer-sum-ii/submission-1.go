func twoSum(numbers []int, target int) (r []int) {
    left := 0
    right := len(numbers) - 1

    for left < right {
        current := numbers[left] + numbers[right]
        if current < target {
            left ++
        } else if current > target {
            right --
        } else {
            return []int{left+1, right+1}
        }
    }
    return r
}
