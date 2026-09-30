func search(nums []int, target int) int {
    left := 0
    right := len(nums) - 1
    for left <= right {
        middle := int((left + right) / 2)
        if nums[middle] == target {
            return middle
        } else if nums[middle] > target {
            right = middle - 1
        } else {
            left = middle + 1
        }
    }
    return -1
}

//   L R
// [-1,0,2,4,6,8]
//   0,1,2,3,4,5
   