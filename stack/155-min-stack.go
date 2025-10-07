package stack

import "math"

type MinStack struct {
	min   int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   math.MaxInt64,
		stack: []int{},
	}
}

func (st *MinStack) Push(val int) {
	if len(st.stack) == 0 {
		st.stack = append(st.stack, 0)
		st.min = val
	} else {
		st.stack = append(st.stack, val-st.min)
		if val < st.min {
			st.min = val
		}
	}
}

func (st *MinStack) Pop() {
	if len(st.stack) == 0 {
		return
	}
	pop := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	if pop < 0 {
		st.min = st.min - pop
	}
}

func (st *MinStack) Top() int {
	top := st.stack[len(st.stack)-1]
	if top > 0 {
		return top + st.min
	}
	return st.min
}

func (st *MinStack) GetMin() int {
	return st.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
