package structure

import (
	"testing"
	"fmt"
	"github.com/bmizerany/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stLen := stack.Len()
	assert.Equal(t, stLen, 4)
	fmt.Printf("stack len %v\n", stLen)
	value := stack.Peak().(int)
	assert.Equal(t, value, 3)
	fmt.Printf("stack Peak value %v\n", value)
	value = stack.Pop().(int)
	assert.Equal(t, value, 3)
	fmt.Printf("stack Pop value %v\n", value)
	afterPopLen := stack.Len()
	assert.Equal(t, afterPopLen, 3)
	fmt.Printf("stack after Pop Len %v\n", value)
	afterPopPeak :=stack.Peak().(int)
	fmt.Printf("stack after Pop Peak %v\n",afterPopPeak)
	isEmpty :=stack.Empty()
	assert.NotEqual(t, isEmpty, true)
	fmt.Printf("now stack isEmpty %v\n",isEmpty)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	nowIsEmpty := stack.Empty()
	assert.Equal(t, nowIsEmpty, true)
	fmt.Printf("now stack isEmpty %v\n",nowIsEmpty)
	nilValue := stack.Pop()
	assert.Equal(t, nilValue, nil)
	fmt.Printf("Empty stack Pop nil value %v\n",nilValue)
}
