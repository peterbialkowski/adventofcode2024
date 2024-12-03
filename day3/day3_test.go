package main

import (
	"testing"
)

func TestProblemOne(t *testing.T) {
	d := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	a := ProblemOne(d)
	if a != 161 {
		t.Error("Expected 161 got", a)
	}
}

func TestProblemTwo(t *testing.T) {
	d := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	a := ProblemTwo(d)
	if a != 48 {
		t.Error("Expected 48 got", a)
	}
}
