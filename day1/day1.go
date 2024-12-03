package main

import (
	"adventofcode2024/cmn"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ProblemOne(left []int, right []int) int {
	// sort both lists, and iterate over each pair adding the sum of the differences and return the total
	var distance int
	sort.Ints(left)
	sort.Ints(right)
	for i, num := range left {
		other := right[i]
		smaller := min(num, other)
		larger := max(num, other)
		distance += larger - smaller
	}
	return distance
}

func ProblemTwo(left []int, right []int) int {
	// count how many times each number in the left list occurrs in the right list. Multiple the number from the left list
	// by how many times it occurs in the right list. Return the sum

	var a int

	for _, num := range left {
		for _, other := range right {
			if num == other {
				a += num
			}
		}
	}

	return a
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	// contents, err := io.ReadAll(file)
	// fmt.Print(lines[0])

	var left []int
	var right []int

	for _, line := range lines {
		fields := strings.Fields(line)
		l, err := strconv.Atoi(fields[0])
		cmn.Check(err)
		r, err := strconv.Atoi(fields[1])
		cmn.Check(err)
		left = append(left, l)
		right = append(right, r)
	}

	answer := ProblemOne(left, right)
	fmt.Println(answer)

	answertwo := ProblemTwo(left, right)
	fmt.Println(answertwo)
}
