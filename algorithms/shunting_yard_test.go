package algos

import "testing"

func Test_shuntingYard_Transform(t *testing.T) {
	tests := []struct {
		name        string
		operators   []string
		infix       string
		wantPostfix string
	}{
		{
			name:        "3+4",
			operators:   []string{"+", "-"},
			infix:       "3+4",
			wantPostfix: "34+",
		},
		{
			name:        "3+4-3",
			operators:   []string{"+", "-"},
			infix:       "3+4-3",
			wantPostfix: "34+3-",
		},
		{
			name:        "3+(4-3)",
			operators:   []string{"+", "-"},
			infix:       "3+(4-3)",
			wantPostfix: "343-+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := NewShuntingYard(tt.operators)
			if gotPostfix := sy.Transform(tt.infix); gotPostfix != tt.wantPostfix {
				t.Errorf("shuntingYard.Transform() = %v, want %v", gotPostfix, tt.wantPostfix)
			}
		})
	}
}
