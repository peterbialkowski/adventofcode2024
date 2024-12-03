package main

import (
	"testing"
)

func TestProblemOne(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := ProblemOne(left, right)
	if result != 11 {
		t.Errorf("ProblemOne = %d; want 11", result)
	}
}

func TestProblemTwo(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := ProblemTwo(left, right)
	if result != 31 {
		t.Errorf("ProblemOne = %d; want 31", result)
	}
}
