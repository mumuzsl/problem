package lco

type MinStack struct {
	A []int
	B []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		A: make([]int, 0),
		B: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.A = append(this.A, x)
	if len(this.B) == 0 || this.B[len(this.B)-1] >= x {
		this.B = append(this.B, x)
	}
}

func (this *MinStack) Pop() {
	var last, temp int
	if last = len(this.A) - 1; last < 0 {
		return
	}
	if temp, this.A = this.A[last], this.A[:last];
		temp == this.B[len(this.B)-1] {
		this.B = this.B[:len(this.B)-1]
	}
}

func (this *MinStack) Top() int {
	return this.A[len(this.A)-1]
}

func (this *MinStack) Min() int {
	return this.B[len(this.B)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
