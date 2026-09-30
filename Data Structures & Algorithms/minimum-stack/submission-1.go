type MinStack struct {
    Stack []int
    Mins []int
}

func Constructor() MinStack {
    return MinStack{Stack: []int{}, Mins: []int{}}
}

func (this *MinStack) Push(val int) {
    if len(this.Mins) == 0 || val <= this.Mins[len(this.Mins)-1]{
        this.Mins = append(this.Mins, val)
    }
    this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
    if this.Mins[len(this.Mins)-1] == this.Stack[len(this.Stack) - 1]{
        this.Mins = this.Mins[:len(this.Mins) - 1]
    }
    this.Stack = this.Stack[:len(this.Stack) - 1]
}

func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack) - 1]
}

func (this *MinStack) GetMin() int {
    return this.Mins[len(this.Mins) - 1]
}
