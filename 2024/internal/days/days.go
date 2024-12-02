package days

import "fmt"

// Function signature for each day's Run function
type DayFunc func()

// Map to register all days dynamically
var dayRegistry = make(map[int]DayFunc)

// RegisterDay allows each day to register itself
func RegisterDay(day int, fn DayFunc) {
	dayRegistry[day] = fn
}

// RunDay runs the solution for a specfic day if it exists
func RunDay(day int) {
	if fn, exists := dayRegistry[day]; exists {
		fn()
	} else {
		fmt.Printf("Day %d not implemented yet!\n", day)
	}
}
