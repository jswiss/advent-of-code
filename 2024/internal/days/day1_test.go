package days

import (
	"2024/utils"
	"fmt"
	"reflect"
	"testing"
)

// Splits an array, returning two new slices of even and odd indexed items
func TestSplitLists(t *testing.T) {
	expectedEvens := []string{"3", "4", "2", "1", "3", "3"}
	expectedOdds := []string{"4", "3", "5", "3", "9", "3"}

	lines, _ := utils.ReadLines("data/day1-example-input.txt")
	evens, odds := SplitLists(lines)

	if reflect.DeepEqual(evens, expectedEvens) {
		t.Log("Evens don't match")
	}
	if reflect.DeepEqual(odds, expectedOdds) {
		t.Log("Odds don't match")
	}
}

func TestOrderList(t *testing.T) {
	input := []string{"3", "4", "2", "1", "3", "3"}
	expected := []int{1, 2, 3, 3, 3, 4}

	result := OrderList(input)

	if reflect.DeepEqual(result, expected) {
		fmt.Println("result and expected do not match")
		fmt.Println("want: {} {}", expected, reflect.TypeOf(expected))
		fmt.Println("got: {} {}", result, reflect.TypeOf(result))
	}
}

func TestRun(t *testing.T) {
	// Example test for Day 1 solution
	t.Log("Day 1 test placeholder")
}
