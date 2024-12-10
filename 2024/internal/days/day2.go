package days

import (
	"2024/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func init() {
	RegisterDay(2, Run)
}

func ProcessString(report string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(report, -1)

	var numbers []int
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		numbers = append(numbers, num)
	}
	return numbers
}

func CheckCompliance(list []int) bool {
	ascending := false
	descending := false
	for i := 1; i < len(list); i++ {
		num1 := list[i-1]
		num2 := list[i]
		diff := math.Abs(float64(num2 - num1))
		ascendingLogic := ((num2 > num1) && (diff >= 1 && diff <= 3))

		fmt.Printf("%d vs %d. Diff is %v. Truthiness is %t.\n\n", num1, num2, diff, ascendingLogic)

		if ascendingLogic {
			ascending = true
		} else {
			ascending = false
			break
		}
	}
	for i := 1; i < len(list); i++ {
		num1 := list[i-1]
		num2 := list[i]
		diff := math.Abs(float64(num2 - num1))
		descendingLogic := ((num2 < num1) && (diff >= 1 && diff <= 3))

		fmt.Printf("%d vs %d. Diff is %v. Truthiness is %t.\n\n", num1, num2, diff, descendingLogic)
		if descendingLogic {
			descending = true
		} else {
			descending = false
			break
		}
	}
	fmt.Printf("Row is %t\n", ascending || descending)
	return ascending || descending
}

func Run() {

	lines, err := utils.ReadLines("data/day2-input.txt")
	if err != nil {
		fmt.Print(err)
	}
	trueCount := 0
	for _, v := range lines {
		line := ProcessString(v)
		fmt.Println(line)
		res := CheckCompliance(line)

		if res {
			trueCount++
		}
	}

	fmt.Println("Solution for Day 1 part 1 is: ", trueCount)
	fmt.Println("Solution for Day 1 part 2 is: ")
}
