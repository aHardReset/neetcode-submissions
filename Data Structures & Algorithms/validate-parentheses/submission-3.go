func isValid(s string) bool {
    openToClose := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    stack := []rune{}
    for _, char := range s {
        if _, isOpen := openToClose[char]; isOpen {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if openToClose[top] != char {
                return false
            }
        }
    }

    return len(stack) == 0
}
