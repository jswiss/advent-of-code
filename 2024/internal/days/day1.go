package days

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func init() {
	RegisterDay(1, Run)
}

func SplitLists(list []string) ([]string, []string) {
	evens := []string{}
	odds := []string{}

	re := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for _, v := range list {
		matches := re.FindStringSubmatch(v)
		if len(matches) > 2 {
			evens = append(evens, matches[1])
			odds = append(odds, matches[2])

		} else {
			fmt.Println("No match found!")
		}
	}
	return evens, odds
}

func OrderList(list []string) []int {
	intList := []int{}
	for _, v := range list {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("issue with strconv")
		}
		intList = append(intList, num)
	}
	sort.Ints(intList)
	return intList
}

func CalculateDifferences(e []int, o []int) []int {
	r := []int{}

	if len(e) != len(o) {
		fmt.Print("Fail!")
	}

	for i := 0; i < len(e); i++ {
		diff := e[i] - o[i]
		if diff < 0 {
			diff = diff * -1
		}
		r = append(r, diff)
	}
	return r
}

func CountOccurrences(e []int, o []int) []int {
	occ := []int{}

	for _, value := range e {
		count := 0
		for _, val := range o {
			if value == val {
				count++
			}
		}
		prod := value * count
		occ = append(occ, prod)
	}
	return occ
}

// Run is the solution for Day 1
// func Run() {
// 	lines, err := utils.ReadLines("data/day1-input.txt")
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	e, o := SplitLists(lines)
// 	evens := OrderList(e)
// 	odds := OrderList(o)

// 	differences := CalculateDifferences(evens, odds)

// 	result := 0

// 	for _, val := range differences {
// 		result += val
// 	}

// 	occ := CountOccurrences(evens, odds)
// 	fmt.Println(occ)

// 	total := 0
// 	for _, v := range occ {
// 		total += v
// 	}
// 	fmt.Println("Solution for Day 1 part 1 is: ", result)
// 	fmt.Println("Solution for Day 1 part 2 is: ", total)
// }
