package main

import "testing"

func TestSimpleCalc(t *testing.T) {
	if Calc(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
func TestCalc(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Test 1", args{x: 5}, 7},
		{"Test 2", args{x: 8}, 10},
		{"Test 3", args{x: -2}, 0},
		{"Test 4", args{x: -9999}, -9997},
		{"Test 5", args{x: 9999}, 10001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Calc(tt.args.x); gotResult != tt.wantResult {
				t.Errorf("Calc() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
