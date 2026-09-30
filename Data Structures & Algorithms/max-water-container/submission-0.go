func maxArea(heights []int) int {
    left := 0
    right := len(heights) - 1

    maxArea := 0

    for left < right {
        area := min(heights[left], heights[right]) * (right - left)
        if area > maxArea {
            maxArea = area
        }
        if heights[left] < heights[right] {
            left++
        } else{
            right--
        }
    }
    return maxArea
}
