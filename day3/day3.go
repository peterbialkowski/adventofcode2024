package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"adventofcode2024/cmn"
)

func ProblemOne(data string) int {
	// r, _ := regexp.Compile(`mul\(\d,\d\)`)
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllString(data, -1)
	// fmt.Println(matches)

	var sum int
	r2, _ := regexp.Compile(`(\d+),(\d+)`)
	fmt.Println("Matches:", len(matches))
	for _, m := range matches {
		digits := r2.FindAllString(m, 2)
		d := strings.Split(digits[0], ",")
		i1, _ := strconv.Atoi(d[0])
		i2, _ := strconv.Atoi(d[1])
		// fmt.Printf("%d * %d = %d", i1, i2, i1*i2)
		// fmt.Println()
		sum += i1 * i2
	}
	return sum
}

func ProblemTwo(data string) int {
	// ignore statements preceeded by don't()
	r, _ := regexp.Compile(`don't\(\).+?do\(\)`)
	data = strings.Replace(data, "\n", "-", -1)
	s := r.ReplaceAllString(data, "x")
	return ProblemOne(s)
}

func main() {
	// read a large string, find instances of mul(\d, \d) and execute a function that multiplies
	// the two integers then return the total.
	content, err := os.ReadFile("input.txt")
	cmn.Check(err)

	answer := ProblemOne(string(content))
	fmt.Println("Problem One:", answer)

	a := ProblemTwo(string(content))
	fmt.Println("Problem Two:", a)
}
