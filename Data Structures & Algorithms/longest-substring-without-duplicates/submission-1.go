func lengthOfLongestSubstring(s string) int {
    repeated := map[rune]struct{}{}
    longest := 0
    left := 0
    right := 0
    for right < len(s) {
        for {
            if _, ok := repeated[rune(s[right])] ; !ok{
                break
            }
            delete(repeated, rune(s[left]))
            left++
        }
        repeated[rune(s[right])] = struct{}{}
        longest = max(longest, len(repeated))  
        right++
    }
    return longest

}


