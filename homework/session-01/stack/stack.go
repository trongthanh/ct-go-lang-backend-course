package stack

type Stack struct {
	data []interface{}
}

// Create a new stack
func New() *Stack {
	stack := Stack{}
	return &stack
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return len(this.data)
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	return this.data[len(this.data)-1]
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	l := len(this.data)
	// get the top element
	top := this.data[l-1]
	// slice off the top element
	this.data = this.data[:l-1]

	return top
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.data = append(this.data, value)
}

