package lco

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	stack1Len := len(this.stack1)
	stack2Len := len(this.stack2)
	if stack2Len > 0 {
		r := this.stack2[stack2Len-1]
		this.stack2 = this.stack2[:stack2Len-1]
		return r
	}

	if stack1Len == 0 {
		return -1
	}

	for i := stack1Len - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	r := this.stack1[0]
	this.stack1 = []int{}

	return r
}
