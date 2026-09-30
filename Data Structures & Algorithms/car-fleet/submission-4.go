import (
    "slices"
)

func carFleet(target int, position []int, speed []int) int {
    posSpeed := sortPosSpeed(position, speed)
    stack := [][2]int{}
    fleets := 0
    for _, current := range posSpeed {
        if len(stack) == 0{
            stack = append(stack, current)
            continue
        }
        top := stack[len(stack)-1]
        if current[1] <= top[1]{
            fleets++
            stack = stack[:len(stack)-1]
            stack = append(stack, current)
            continue
        }
        intersection := float64(top[0] - current[0]) / float64(current[1] - top[1])
        point := float64(top[0]) + (intersection * float64(top[1]))
        if point > float64(target) {
            stack = append(stack, current)
        }
    }
    return fleets + len(stack)
}

func sortPosSpeed(position, speed []int) [][2]int{
    posSpeed := [][2]int{}
    for i, pos := range position {
        posSpeed = append(posSpeed, [2]int{pos, speed[i]})
    }
    slices.SortFunc(posSpeed, func(a ,b [2]int) int {
        return b[0] - a[0]
    })
    return posSpeed
}
