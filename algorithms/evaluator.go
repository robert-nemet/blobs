package algorithms

import (
	"strings"

	"github.com/robert-nemet/blobs/datas"
)

type Operation[T any] func(stack datas.Stack[T]) (T, error)
type Convert[T any] func(input string) (T, error)

type Evaluator[T any] interface {
	Evaluate(input string) (T, error)

	isOperation(token string) bool
}

type evaluator[T any] struct {
	operators map[string]Operation[T]
	convert   Convert[T]
}

func NewEvaluator[T any](operators map[string]Operation[T], convert Convert[T]) Evaluator[T] {
	return evaluator[T]{
		operators: operators,
		convert:   convert,
	}
}

func (e evaluator[T]) Evaluate(postfix string) (T, error) {
	stack := datas.NewStack[T](nil)
	input := strings.Split(postfix, " ")
	for _, token := range input {
		if e.isOperation(token) {
			newop, err := e.operators[token](stack)
			if err != nil {
				return *new(T), err
			}
			stack.Push(newop)
		} else {
			val, err := e.convert(token)
			if err != nil {
				return *new(T), err
			}
			stack.Push(val)
		}
	}

	return *stack.Pop(), nil
}

func (e evaluator[T]) isOperation(token string) bool {
	for key := range e.operators {
		if key == token {
			return true
		}
	}

	return false
}
