func productExceptSelf(nums []int) []int {
    var zeros uint8
    allExceptZeros := 1

    for _, n := range nums {
        if n == 0 {
            zeros++
        } else {
            if zeros <= 1 {
                allExceptZeros *= n
            } else {
                allExceptZeros = 0
                break
            }
        }
    }

    res := make([]int, len(nums))

    if allExceptZeros == 0 || zeros > 1 {
        return res[:len(nums)]
    } 
    for i, n := range nums{
        if zeros == 1{
            if n == 0 {
                res[i] = allExceptZeros
            }
        } else {
            res[i] = int(allExceptZeros / n)
        }
        
    }
    return res[:len(nums)]
}
