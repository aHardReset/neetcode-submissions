
func dailyTemperatures(temperatures []int) []int {
    stack := [][]int{}
    res := make([]int, len(temperatures))

    for i := len(temperatures) - 1; i >= 0; i-- {

        for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1][0] {
            stack = stack[:len(stack)-1]
        }

        if len(stack) == 0 {
            res[i] = 0
        } else {
            res[i] = stack[len(stack)-1][1] - i
        }
        stack = append(stack, []int{temperatures[i], i})
    }

    return res

}

// [30,38,30,36,35,40,28] 6
// [40 5, 38 1, 30 0]  

// [1,4,1,2,1,0,0]