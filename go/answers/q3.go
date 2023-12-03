package answers

import (
	"strconv"
)

func Day3() []interface{} {
	data := ReadInputAsStr(3)
	return []interface{}{q3part1(data), q3part2(data)}
}

func integerNextToSymbol(x int, y int, symbol_coordinates [][]int) bool {
	for _, coords := range symbol_coordinates {
		xdiff := x - coords[0]
		ydiff := y - coords[1]
		if abs(xdiff) <= 1 && abs(ydiff) <= 1 {
			return true
		}
	}
	return false
}

func q3part1(data []string) int {
	not_symbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	symbol_coordinates := [][]int{}
	for y, row := range data {
		for x, cell := range row {
			// Anything that isn't a digit or a . is a symbol
			if RuneNotInSlice(cell, not_symbols) {
				symbol_coordinates = append(symbol_coordinates, []int{x, y})
			}
		}
	}

	total := 0
	for y, row := range data {
		active_number := ""
		for x, cell := range row {
			if RuneInSlice(cell, digits) {
				active_number += string(cell)
				continue
			}

			if active_number != "" {
				is_adjacent := false
				active_int, _ := strconv.Atoi(active_number)
				for offset := 1; offset <= len(active_number); offset++ {
					if integerNextToSymbol(x-offset, y, symbol_coordinates) {
						is_adjacent = true
						break
					}
				}
				if is_adjacent {
					total += active_int
				}
				active_number = ""
			}

		}
		if active_number != "" {
			is_adjacent := false
			active_int, _ := strconv.Atoi(active_number)
			for offset := 0; offset <= len(active_number); offset++ {
				if integerNextToSymbol(len(row)-offset, y, symbol_coordinates) {
					is_adjacent = true
					break
				}
			}
			if is_adjacent {
				total += active_int
			}
			active_number = ""
		}
	}

	return total
}

func find_adjacent_gear(x int, y int, num_length int, x_max int, value int, gearmap map[int]int) int {
	for x_new := x - num_length - 1; x_new <= x; x_new++ {
		for y_new := y - 1; y_new <= y+1; y_new++ {
			pos := y_new*x_max + x_new
			other_value, exists := gearmap[pos]
			if exists == false {
				continue
			}
			if other_value == 0 {
				gearmap[pos] = value
			} else if other_value == value {
				continue
			} else {
				result := other_value * value
				delete(gearmap, pos)
				return result
			}
		}
	}
	return 0
}

func q3part2(data []string) int {
	x_max := len(data[0])
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	gear_map := map[int]int{}
	for y, row := range data {
		for x, cell := range row {
			// Anything that isn't a digit or a . is a symbol
			if cell == '*' {
				pos := y*x_max + x
				gear_map[pos] = 0
			}
		}
	}

	total := 0
	for y, row := range data {
		active_number := ""
		for x, cell := range row {
			if RuneInSlice(cell, digits) {
				active_number += string(cell)
				continue
			}

			if active_number != "" {
				active_int, _ := strconv.Atoi(active_number)
				total += find_adjacent_gear(x, y, len(active_number), x_max, active_int, gear_map)
				active_number = ""
			}
		}
		if active_number != "" {
			active_int, _ := strconv.Atoi(active_number)
			total += find_adjacent_gear(x_max, y, len(active_number), x_max, active_int, gear_map)
			active_number = ""
		}
	}

	return total
}

// 544502 Too high
// 546817 Too high
// 523575 Wrong
// 536676 Wrong
