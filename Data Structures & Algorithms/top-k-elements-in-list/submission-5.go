
func topKFrequent(nums []int, k int) []int {
    c := map[int]int{}

    for _, num := range nums{
        c[num]++
    }

    freqSlice := make([][]int, len(nums) + 1)
    
    for num, freq := range c {
        freqSlice[freq] = append(freqSlice[freq], num)
    }

    res := []int{}
    for i := len(freqSlice) -1 ; i >= 0 ; i-- {
        for _, num := range freqSlice[i] {
            res = append(res, num)
        }
        if len(res) >= k {
            break
        }
    }
    return res[:k]
    
}
