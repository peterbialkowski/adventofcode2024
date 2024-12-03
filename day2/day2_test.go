package main

import (
	"fmt"
	"testing"
)

func TestProblemOne(t *testing.T) {
	r1 := Report{
		levels: []int{7, 6, 4, 2, 1},
	}
	a1 := Safe(r1)
	if a1 != true {
		t.Error("ProblemOne = false, wanted true")
	}

	r2 := Report{
		levels: []int{1, 3, 6, 7, 9},
	}
	a2 := Safe(r2)
	if a2 != true {
		t.Error(r2.levels, "should have been true")
	}

	r3 := Report{
		levels: []int{1, 2, 7, 8, 9},
	}
	a3 := Safe(r3)
	if a3 != false {
		t.Error(r3.levels, "should have been false")
	}
}

func TestProblemTwo(t *testing.T) {
	reports := []Report{
		{levels: []int{7, 6, 4, 2, 1}},
		{levels: []int{1, 2, 7, 8, 9}},
		{levels: []int{9, 7, 6, 2, 1}},
		{levels: []int{1, 3, 2, 4, 5}},
		{levels: []int{8, 6, 4, 4, 1}},
		{levels: []int{1, 3, 6, 7, 9}},
	}

	totalSafe := 0

	for i, r := range reports {
		fmt.Println("Testing Report", i)
		a := SafeTwo(r)
		if a == true {
			totalSafe += 1
		}
	}
	if totalSafe != 4 {
		t.Error("Should have been 4 not", totalSafe)
	}
}
