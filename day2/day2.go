package main

import (
	"bufio"
	"fmt"
	// "fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type data struct {
	reports []Report
}

type Report struct {
	levels []int
}

func remove(slice []int, s int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:s]...)
	return append(ret, slice[s+1:]...)
}

func SafeTwo(report Report) bool {
	// iterate over slice removing one element at a time
	// if it is safe then return true, if all tests fail return false
	for i := range report.levels {
		fmt.Println("remove element", i)
		subslice := remove(report.levels, i)
		fmt.Println("Checking", subslice)
		report := Report{
			levels: subslice,
		}
		safe := Safe(report)
		if safe == true {
			return true
		}
	}
	return false

}

func Safe(report Report) bool {
	// report.levels is a slice of ints
	// it is safe if the levels are either all increasing or all decreasing
	// AND, two adjacent levels differ by at least one and at most three

	// first test if it's already increasing or decreasing
	increasing := sort.IntsAreSorted(report.levels)
	d := make([]int, len(report.levels))
	_ = copy(d, report.levels)
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	decreasing := reflect.DeepEqual(d, report.levels)
	// fmt.Println("d:", d, "; report.levels:", report.levels)
	if !decreasing && !increasing {
		// fmt.Println(report.levels)
		// println("increasing: %s; decreasing %s", increasing, decreasing)
		return false
	}

	for i, num := range report.levels {
		if i == len(report.levels)-1 {
			break
		}
		distance := num - report.levels[i+1]
		if distance == 0 {
			return false
		}

		if distance > 3 || distance < -3 {
			return false
		}

	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	var reports []Report
	for sc.Scan() {
		rs := strings.Fields(sc.Text())
		ints := make([]int, len(rs))
		for i, s := range rs {
			ints[i], _ = strconv.Atoi(s)
		}
		reports = append(reports, Report{levels: ints})
	}

	safe := 0

	for _, report := range reports {
		isSafe := Safe(report)
		if isSafe == true {
			safe += 1
		}
	}

	println(safe)

	safe2 := 0

	for _, report := range reports {
		isSafe := SafeTwo(report)
		if isSafe == true {
			safe2 += 1
		}
	}

	println(safe2)
}
