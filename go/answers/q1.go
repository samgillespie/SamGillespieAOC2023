package answers

import (
	"strconv"
	"strings"
)

func Day1() []interface{} {
	data := ReadInputAsStr(1)
	return []interface{}{q1part1(data), q1part2(data)}

}

func extract_digits(input string) int {
	digit1 := -1
	digit2 := -1
	digit1_set := false
	var err error
	for _, character := range input {
		if digit1_set == false {
			digit1, err = strconv.Atoi(string(character))
			if err == nil {
				digit1_set = true
			}
		} else {
			digit2temp, err := strconv.Atoi(string(character))
			if err == nil {
				digit2 = digit2temp
			}
		}
	}
	if digit2 == -1 {
		digit2 = digit1
	}
	return digit1*10 + digit2
}

func find_all_matches(input string, substring string, number int) []int {
	match_locations := []int{}
	step := 0
	initial := true
	for {
		if initial == false {
			step += 1
		}
		text_position := strings.Index(input[step:], substring)
		num_position := strings.Index(input[step:], strconv.Itoa(number))
		// fmt.Println(input, substring, step, text_position, num_position)
		if text_position == -1 && num_position == -1 {
			return match_locations
		}

		if text_position != -1 {
			if intSliceContains(match_locations, text_position+step) == false {
				match_locations = append(match_locations, text_position+step)
			}
		}
		if num_position != -1 {
			if intSliceContains(match_locations, num_position+step) == false {
				match_locations = append(match_locations, num_position+step)
			}
		}
		step += min_exclude_minus(num_position, text_position)
		initial = false
	}
}

func extract_numbers(input string) int {
	lookup := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	max_position := -1
	min_position := 99999
	max_value := -1
	min_value := -1
	for text_number, value := range lookup {
		matched_positions := find_all_matches(input, text_number, value)
		for _, position := range matched_positions {
			if position > max_position {
				max_position = position
				max_value = value
			}
			if position < min_position {
				min_position = position
				min_value = value
			}
		}
		// fmt.Println(text_number, value, max_position, min_position, min_value*10+max_value, matched_positions)
	}
	// fmt.Println(input, min_position, max_position, min_value*10+max_value)
	if min_value < 0 || max_value < 0 {
		panic("h")
	}
	return min_value*10 + max_value
}

func q1part1(data []string) int {
	sum := 0
	for _, elem := range data {
		value := extract_digits(elem)
		sum += value
	}
	return sum
}

func q1part2(data []string) int {
	sum := 0
	for _, elem := range data {
		value := extract_numbers(elem)
		// fmt.Println(elem, value)
		sum += value
	}
	return sum
}

// Wrong  54254
