package algos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stacks(t *testing.T) {
	one := "one"
	two := "two"
	three := "three"
	four := "four"

	stack := NewStack(&one)
	assert.Equal(t, one, *stack.Peek())

	stack.Push(two)
	stack.Push(three)

	assert.Equal(t, three, *stack.Peek())

	assert.Equal(t, three, *stack.Pop())
	assert.Equal(t, two, *stack.Peek())

	assert.Equal(t, two, *stack.Pop())
	assert.Equal(t, one, *stack.Peek())

	stack.Push(three)

	assert.Equal(t, three, *stack.Pop())
	assert.Equal(t, one, *stack.Peek())

	stack.Pop()

	assert.Nil(t, stack.Peek())
	assert.Nil(t, stack.Pop())
	assert.True(t, stack.IsEmpty())

	stack.Push(four)

	assert.Equal(t, four, *stack.Peek())
}

func Test_StackNil(t *testing.T) {
	stack := NewStack[string](nil)

	assert.Nil(t, stack.Peek())
	assert.Nil(t, stack.Pop())
	assert.True(t, stack.IsEmpty())

	stack.Push("one")

	assert.Equal(t, "one", *stack.Peek())
	assert.Equal(t, "one", *stack.Pop())
}

func Test_Flush(t *testing.T) {
	stack := NewStack[string](nil)

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	assert.Equal(t, []string{"three", "two", "one"}, stack.Flush())
	assert.Equal(t, []string{}, stack.FlushReverse())
}

func Test_FlushReverse(t *testing.T) {
	stack := NewStack[string](nil)

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	assert.Equal(t, []string{"one", "two", "three"}, stack.FlushReverse())
}
