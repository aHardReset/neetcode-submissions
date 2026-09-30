func evalRPN(tokens []string) int {
    stack := []int{}

    for _, t := range tokens {
        num, err := strconv.Atoi(t)
        if err == nil {
            stack = append(stack, num)
        } else {
            reduceStack(&stack, t)
        }
    }

    return stack[0]
}

func reduceStack(stack *[]int, operation string) {
    b := (*stack)[len(*stack) - 1]
    a := (*stack)[len(*stack) - 2]
    (*stack) = (*stack)[:len(*stack) - 2]
    var r int
    switch operation{
        case "+":
            r = a + b
        case "-":
            r = a - b
        case "*":
            r = a * b
        case "/":
            r = int(a/b)
        default:
            r = 0
    }
    (*stack) = append(*stack, r)
}
