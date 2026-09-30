
func threeSum(nums []int) [][]int {
    res := [][]int{}
    // nums = removeDuplicates(nums)
    sort.Ints(nums)
    
    for i := 0; i < len(nums); i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        results := twoSum(nums[i+1:], -nums[i])
        for _, r := range results{
            res = append(res,[]int{nums[i], r[0], r[1]})
        }

    }
    return res
}

func twoSum(nums []int, target int) [][]int {
    l := 0
    r := len(nums) - 1
    res := [][]int{}

    for l < r {
        current := nums[l] + nums[r]
        if current == target {
            res = append(res, []int{nums[l], nums[r]})
            l ++
            for nums[l] == nums[l-1] && l < r{
                l++
            }
        } else if current > target{
            r--
        } else {
            l ++
        }
    }
    return res
}

// [-1,0,1,2,-1,-4]
// [-4,-1,0,1,2]
//      l   r
// 0
// 