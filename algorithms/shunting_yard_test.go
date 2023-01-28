package algos

import "testing"

func Test_shuntingYard_Transform(t *testing.T) {
	tests := []struct {
		name        string
		operators   map[string]int
		infix       string
		wantPostfix string
	}{
		{
			name:        "3 + 4",
			operators:   map[string]int{"+": 1, "-": 1},
			infix:       "3 + 4",
			wantPostfix: "3 4 +",
		},
		{
			name:        "3 + 4 - 3",
			operators:   map[string]int{"+": 1, "-": 1},
			infix:       "3 + 4 - 3",
			wantPostfix: "3 4 + 3 -",
		},
		{
			name:        "3 + ( 4 - 3 )",
			operators:   map[string]int{"+": 1, "-": 1},
			infix:       "3 + ( 4 - 3 )",
			wantPostfix: "3 4 3 - +",
		},
		{
			name:        "( (A * B) + (C / D) )",
			operators:   map[string]int{"+": 1, "-": 1, "*": 2, "/": 2},
			infix:       "( ( A * B ) + ( C / D ) )",
			wantPostfix: "A B * C D / +",
		},
		{
			name:        "((A * (B + C) ) / D)",
			operators:   map[string]int{"+": 1, "-": 1, "*": 2, "/": 2},
			infix:       "( ( A * ( B + C ) ) / D )",
			wantPostfix: "A B C + * D /",
		},
		{
			name:        "(A * (B + (C / D) ) )",
			operators:   map[string]int{"+": 1, "-": 1, "*": 2, "/": 2},
			infix:       "( A * ( B + ( C / D ) ) )",
			wantPostfix: "A B C D / + *",
		},
		{
			name:        "B + C / D + B * A",
			operators:   map[string]int{"+": 1, "-": 1, "*": 2, "/": 2},
			infix:       "B + C / D + B * A",
			wantPostfix: "B C D / + B A * +",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := NewShuntingYard(tt.operators)
			if gotPostfix := sy.Transform(tt.infix); gotPostfix != tt.wantPostfix {
				t.Errorf("shuntingYard.Transform(%s) = %v, want %v", tt.infix, gotPostfix, tt.wantPostfix)
			}
		})
	}
}
