package algorithms

import (
	"errors"
	"strconv"
	"testing"

	"github.com/robert-nemet/blobs/datas"
	"github.com/stretchr/testify/assert"
)

func Test_EvaluateExpressionInt(t *testing.T) {

	operations := map[string]Operation[int]{
		"+": func(s datas.Stack[int]) (int, error) {
			fop := s.Pop()
			if fop == nil {
				return 0, errors.New("no operand")
			}
			sop := s.Pop()
			if sop == nil {
				return 0, errors.New("no operand")
			}
			return *fop + *sop, nil
		},
		"-": func(s datas.Stack[int]) (int, error) {
			sop := s.Pop()
			if sop == nil {
				return 0, errors.New("no operand")
			}
			fop := s.Pop()
			if sop == nil {
				return 0, errors.New("no operand")
			}
			return *fop - *sop, nil
		},
	}
	evaluator := NewEvaluator(operations, func(input string) (int, error) {
		return strconv.Atoi(input)
	})

	tests := []struct {
		name      string
		expresion string
		expect    int
		err       bool
	}{
		{
			name:      "3 4 +",
			expresion: "3 4 +",
			expect:    7,
			err:       false,
		},
		{
			name:      "3 4 + 3 -",
			expresion: "3 4 + 3 -",
			expect:    4,
			err:       false,
		},
		{
			name:      "3 4 3 - +",
			expresion: "3 4 3 - +",
			expect:    4,
			err:       false,
		},
		{
			name:      "3 4 + f",
			expresion: "3 4 + f",
			expect:    7,
			err:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := evaluator.Evaluate(tt.expresion)

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expect, res)
			}
		})
	}

}

func Test_EvaluateExpressionBool(t *testing.T) {

	operations := map[string]Operation[bool]{
		"&": func(s datas.Stack[bool]) (bool, error) {
			sop := s.Pop()
			if sop == nil {
				return false, errors.New("no operand")
			}
			fop := s.Pop()
			if fop == nil {
				return false, errors.New("no operand")
			}
			return *fop && *sop, nil
		},
		"|": func(s datas.Stack[bool]) (bool, error) {
			sop := s.Pop()
			if sop == nil {
				return false, errors.New("no operand")
			}
			fop := s.Pop()
			if fop == nil {
				return false, errors.New("no operand")
			}
			return *fop || *sop, nil
		},
		"^": func(s datas.Stack[bool]) (bool, error) {
			op := s.Pop()
			if op == nil {
				return false, errors.New("no operand")
			}
			return !*op, nil
		},
	}
	evaluator := NewEvaluator(operations, func(input string) (bool, error) {
		return strconv.ParseBool(input)
	})

	tests := []struct {
		name      string
		expresion string
		expect    bool
		err       bool
	}{
		{
			name:      "true true ^ &",
			expresion: "true true ^ &",
			expect:    false,
			err:       false,
		},
		{
			name:      "true true &",
			expresion: "true true &",
			expect:    true,
			err:       false,
		},
		{
			name:      "true false | false &",
			expresion: "true false | false &",
			expect:    false,
			err:       false,
		},
		{
			name:      "true false true | &",
			expresion: "true false true | &",
			expect:    true,
			err:       false,
		},
		{
			name:      "true false true | & xxx",
			expresion: "true false true | & xxx",
			expect:    false,
			err:       true,
		},
		{
			name:      "false ^ false ^ &",
			expresion: "false ^ false ^ &",
			expect:    true,
			err:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := evaluator.Evaluate(tt.expresion)

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expect, res)
			}
		})
	}
}
