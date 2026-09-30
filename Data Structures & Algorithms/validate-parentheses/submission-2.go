func isValid(s string) bool {
    openToClose := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    stack := []rune{}
    r := []rune(s)
    for i := 0 ; i < len(r) ; i++ {
        if _, isOpen := openToClose[r[i]]; isOpen {
            stack = append(stack, r[i])
        } else {
            if top, ok := pop(&stack); ok{
                topClose, _ := openToClose[top]
                if topClose != r[i] {
                    return false // closed not matched
                }
            } else {
                return false // closed when no left in stack
            }
        }
    }

    return len(stack) == 0
}

func pop(stack *[]rune) (rune, bool) {
    var top rune

    if len(*stack) == 0 {
        return top, false
    }

    top = (*stack)[len(*stack)-1]
    (*stack) = (*stack)[:len(*stack)-1]
    return top, true
}
