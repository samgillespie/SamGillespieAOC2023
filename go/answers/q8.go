package answers

import (
	"fmt"
	"strings"
)

func Day8() []interface{} {
	data := ReadInputAsStr(8)
	instructions := data[0]
	paths := map[string][]string{}
	for _, row := range data[2:] {
		row = strings.ReplaceAll(row, "(", "")
		row = strings.ReplaceAll(row, ")", "")
		split := strings.Split(row, " = ")
		split_paths := strings.Split(split[1], ", ")
		paths[split[0]] = split_paths
	}
	return []interface{}{q8part1(instructions, paths), q8part2(instructions, paths)}
}

func q8part1(instructions string, paths map[string][]string) int {
	position := paths["AAA"]
	steps_taken := 0
	for {
		for _, instruction_rune := range instructions {
			var next_path string
			if instruction_rune == 'L' {
				next_path = position[0]
			} else {
				next_path = position[1]
			}
			position = paths[next_path]
			steps_taken++
			if next_path == "ZZZ" {
				return steps_taken
			}
		}
	}
}

func multiplyAllListInts(input []int) int {
	value := 1
	for _, i := range input {
		value = value * i
	}
	return value
}

func q8part2(instructions string, paths map[string][]string) int {
	positions := []string{}
	first_time_reached := []int{}
	for position := range paths {
		if position[2] == 'A' {
			positions = append(positions, position)
			first_time_reached = append(first_time_reached, 0)
		}
	}

	fmt.Println("Starting at positions ", positions)
	steps_taken := 0
	counter := 0
	for {
		for _, instruction_rune := range instructions {
			for i, position := range positions {
				var next_path string
				options := paths[position]
				if instruction_rune == 'L' {
					next_path = options[0]
				} else {
					next_path = options[1]
				}
				positions[i] = next_path
			}

			steps_taken++
			// Check if finished
			success := true
			for i, position := range positions {
				if position[2] == 'Z' {

					if first_time_reached[i] == 0 {
						first_time_reached[i] = steps_taken
						counter++
					}
					fmt.Println(steps_taken, i, first_time_reached)
				} else {
					success = false
				}
			}
			if counter == len(positions) {
				lcm := 1
				for i := range first_time_reached {
					lcm = LCM(lcm, first_time_reached[i])
				}
				return lcm
			}
			if success == true {
				return steps_taken
			}
		}
	}
}
