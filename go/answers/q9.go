package answers

import (
	"strings"
)

func Day9() []interface{} {
	data := ReadInputAsStr(9)
	sequences := [][]int{}
	for _, row := range data {
		split_str := strings.Split(row, " ")
		sequences = append(sequences, toListOfInts(split_str))
	}
	return []interface{}{q9part1(sequences), q9part2(sequences)}
}

func dydx(sequence []int) []int {
	derivative := make([]int, len(sequence)-1)
	for i := range sequence {
		if i == len(sequence)-1 {
			// Skip last element
			continue
		}
		difference := sequence[i+1] - sequence[i]
		derivative[i] = difference
	}
	return derivative
}

func integral(sequence []int, initial_value int) []int {
	derivative := make([]int, len(sequence)-1)
	for i := range sequence {
		if i == len(sequence)-1 {
			// Skip last element
			continue
		}
		difference := sequence[i+1] - sequence[i]
		derivative[i] = difference
	}
	return derivative
}

func is_sequence_consistent(sequence []int) bool {
	// Returns true of all elements are the same
	var value int
	for i, elem := range sequence {
		if i == 0 {
			value = elem
			continue
		}
		if elem != value {
			return false
		}
	}
	return true
}

func q9part1(data [][]int) int {
	next_steps := make([]int, len(data))
	for sequence_num, sequence := range data {
		order := 0
		initial_values := []int{sequence[0]}
		derivative := sequence
		for {
			derivative = dydx(derivative)
			initial_values = append(initial_values, derivative[0])
			order++

			if is_sequence_consistent(derivative) {
				// Add one more so we can see the next number
				derivative = append(derivative, derivative[0])
				break
			}

		}
		// Regenerate the sequence
		for i := 0; i < order; i++ {
			next_sequence := make([]int, len(derivative)+1)
			index := order - i - 1
			next_sequence[0] = initial_values[index]
			for j := 1; j < len(derivative)+1; j++ {
				next_sequence[j] = next_sequence[j-1] + derivative[j-1]
			}
			derivative = next_sequence
		}
		next_steps[sequence_num] = derivative[len(derivative)-1]
	}
	return sumSlice(next_steps)
}

func q9part2(data [][]int) int {
	prev_steps := make([]int, len(data))
	for sequence_num, sequence := range data {
		order := 0
		initial_values := []int{sequence[0]}
		derivative := sequence
		var rate int
		for {
			derivative = dydx(derivative)
			initial_values = append(initial_values, derivative[0])
			order++

			if is_sequence_consistent(derivative) {
				rate = derivative[0]
				break
			}
		}
		// Regenerate the sequence
		for i := 0; i < order; i++ {
			index := order - i - 1
			first_value := initial_values[index]
			rate = first_value - rate
		}
		prev_steps[sequence_num] = rate
	}
	return sumSlice(prev_steps)
}
