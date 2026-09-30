func isPalindrome(s string) bool {
    sLower := strings.ToLower(s)
    left := 0
    right := len(sLower) -1
    
    for left < right {
        for left < right {
            if isAlphanumeric(sLower[left]){
                break
            } else {
                left++
            }
        }

        for right > left {
            if isAlphanumeric(sLower[right]){
                break
            } else {
                right--
            }
        }

        if sLower[left] != sLower[right] {
            return false
        }
        left ++
        right--
    }

    return true
}

func isAlphanumeric(r uint8) bool {
    return (r >= 97 && r <=122) || (r >= 48 && r<=57)
}