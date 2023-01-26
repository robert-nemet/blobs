package algos

import "strings"

/*
 * Convert expresion from infix notation to postfix notation
 */

type shuntingYard struct {
	Operators      []string
	OutputStack    Stack[string]
	OperatorsStack Stack[string]
}

type ShuntingYard interface {
	Transform(infix string) (postfix string)

	isOp(op rune) bool
}

func NewShuntingYard(oprs []string) ShuntingYard {
	return shuntingYard{
		Operators:      oprs,
		OutputStack:    NewStack[string](nil),
		OperatorsStack: NewStack[string](nil),
	}
}

func (sy shuntingYard) Transform(infix string) (postfix string) {
	for _, s := range infix {
		v := string(s)
		if sy.isOp(s) {
			if sy.OperatorsStack.IsEmpty() {
				sy.OperatorsStack.Push(v)
			} else {
				oldOp := *sy.OperatorsStack.Pop()
				sy.OutputStack.Push(oldOp)
				sy.OperatorsStack.Push(v)
			}
		} else {
			sy.OutputStack.Push(v)
		}
	}

	for !sy.OperatorsStack.IsEmpty() {
		sy.OutputStack.Push(*sy.OperatorsStack.Pop())
	}

	fv := sy.OutputStack.FlushReverse()
	// for i, j := 0, len(fv)-1; i < j; i, j = i+1, j-1 {
	// 	fv[i], fv[j] = fv[j], fv[i]
	// }
	return strings.Join(fv, "")
}

func (sy shuntingYard) isOp(op rune) bool {
	ops := string(op)
	for _, opv := range sy.Operators {
		if ops == opv {
			return true
		}
	}
	return false
}
