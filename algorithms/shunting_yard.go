package algos

import (
	"strings"
)

/*
 * Convert expresion from infix notation to postfix notation
 */

type shuntingYard struct {
	Operators      map[string]int
	OutputStack    Stack[string]
	OperatorsStack Stack[string]
}

type ShuntingYard interface {
	Transform(infix string) (postfix string)

	isOp(op string) bool
	isValue(op string) bool
	isBracket(op string) bool
	isLeftBracket(op string) bool
	isRightBracket(op string) bool
	greater(newOp, oldOp string) bool
}

func NewShuntingYard(oprs map[string]int) ShuntingYard {
	return shuntingYard{
		Operators:      oprs,
		OutputStack:    NewStack[string](nil),
		OperatorsStack: NewStack[string](nil),
	}
}

func (sy shuntingYard) Transform(infix string) (postfix string) {
	infixArray := strings.Split(infix, " ")
	for _, v := range infixArray {
		switch {
		case sy.isValue(v):
			sy.OutputStack.Push(v)
		case sy.isLeftBracket(v):
			sy.OperatorsStack.Push(v)
		case sy.isRightBracket(v):
			for {
				if sy.OperatorsStack.IsEmpty() {
					break
				}
				nv := sy.OperatorsStack.Pop()
				if sy.isLeftBracket(*nv) {
					break
				}
				sy.OutputStack.Push(*nv)
			}
		default:
			if !sy.OperatorsStack.IsEmpty() && !sy.isLeftBracket(*sy.OperatorsStack.Peek()) {
				for !sy.OperatorsStack.IsEmpty() && !sy.greater(v, *sy.OperatorsStack.Peek()) {
					oldOp := *sy.OperatorsStack.Pop()
					sy.OutputStack.Push(oldOp)
				}
			}
			sy.OperatorsStack.Push(v)
		}
	}

	for !sy.OperatorsStack.IsEmpty() {
		sy.OutputStack.Push(*sy.OperatorsStack.Pop())
	}

	fv := sy.OutputStack.FlushReverse()
	return strings.Join(fv, " ")
}

func (sy shuntingYard) isOp(op string) bool {
	for k, _ := range sy.Operators {
		if op == k {
			return true
		}
	}
	return false
}

func (sy shuntingYard) isValue(op string) bool {
	return !(sy.isOp(op) || sy.isBracket(op))
}

func (sy shuntingYard) isBracket(op string) bool {
	return sy.isLeftBracket(op) || sy.isRightBracket(op)
}

func (sy shuntingYard) isLeftBracket(op string) bool {
	return op == "("
}

func (sy shuntingYard) isRightBracket(op string) bool {
	return op == ")"
}

func (sy shuntingYard) greater(newOp, oldOp string) bool {
	return sy.Operators[newOp] > sy.Operators[oldOp]
}
