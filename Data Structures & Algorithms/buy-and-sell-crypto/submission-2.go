func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    left := 0
    right := 1
    maxVal := 0

    for right < len(prices) {
        if prices[left] >= prices[right] {
            left = right
            right++
        } else{
            maxVal = max(maxVal, prices[right] - prices[left])
            right++
        }
    }
    return maxVal
}
