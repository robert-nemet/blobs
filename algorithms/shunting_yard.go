package algorithms

import (
	"strings"

	"github.com/robert-nemet/blobs/datas"
)

/*
 * Convert expresion from infix notation to postfix notation
 */

type shuntingYard struct {
	Operators map[string]int
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
		Operators: oprs,
	}
}

func (sy shuntingYard) Transform(infix string) (postfix string) {
	infixArray := strings.Split(infix, " ")
	out := datas.NewStack[string](nil)
	ops := datas.NewStack[string](nil)

	for _, v := range infixArray {
		switch {
		case sy.isValue(v):
			out.Push(v)
		case sy.isLeftBracket(v):
			ops.Push(v)
		case sy.isRightBracket(v):
			for {
				if ops.IsEmpty() {
					break
				}
				nv := ops.Pop()
				if sy.isLeftBracket(*nv) {
					break
				}
				out.Push(*nv)
			}
		default:
			if !ops.IsEmpty() && !sy.isLeftBracket(*ops.Peek()) {
				for !ops.IsEmpty() && !sy.greater(v, *ops.Peek()) {
					oldOp := *ops.Pop()
					out.Push(oldOp)
				}
			}
			ops.Push(v)
		}
	}

	for !ops.IsEmpty() {
		out.Push(*ops.Pop())
	}

	fv := out.FlushReverse()
	return strings.Join(fv, " ")
}

func (sy shuntingYard) isOp(op string) bool {
	for k := range sy.Operators {
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
