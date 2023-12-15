package answers

import (
	"fmt"
	"math"
	"strings"
)

func Day15() []interface{} {
	data := ReadInputAsStr(15)
	input_codes := strings.Split(data[0], ",")
	return []interface{}{q15part1(input_codes), q15part2(input_codes)}
}

func q15part1(input_codes []string) int {
	total := 0
	for _, code := range input_codes {
		value := HASH_ALGO(code)
		total += value
	}
	return total
}

func box_score(boxes [][]string) int {
	score := 0
	for idx, box := range boxes {
		for slot, elem := range box {
			focal_length := int(elem[len(elem)-1] - 48)
			score += (idx + 1) * (slot + 1) * focal_length
		}
	}
	return score
}

func print_boxes(boxes [][]string) {
	for idx, box := range boxes {
		if len(box) == 0 {
			continue
		}
		fmt.Printf("box: %d, %s|", idx, box)
	}
	fmt.Println()
}

func q15part2(input_codes []string) int {
	boxes := make([][]string, 256)
	for _, code := range input_codes {
		is_minus := StringContainsRune(code, '-')
		var subcode string
		if is_minus {
			subcode = code[0 : len(code)-1]
		} else {
			subcode = code[0 : len(code)-2]
		}
		box_num := HASH_ALGO(subcode)
		box := boxes[box_num]

		if is_minus {
			for pos, elem := range box {
				if len(subcode) > len(elem) {
					continue
				}

				if elem[0:len(elem)-2] == subcode {
					box = append(box[0:pos], box[pos+1:]...)
					break
				}
			}
			boxes[box_num] = box
		} else {
			// = logic
			replaced := false
			for pos, elem := range box {
				if len(subcode) > len(elem) {
					continue
				}
				if elem[0:len(elem)-2] == subcode {
					box[pos] = code
					replaced = true
					break
				}
			}
			if replaced == false {
				box = append(box, code)
			}
			boxes[box_num] = box
		}
		// fmt.Println(code, subcode)
		// print_boxes(boxes)
	}
	return box_score(boxes)
}

func HASH_ALGO(input string) int {
	value := 0
	for _, character := range []rune(input) {
		value += int(character)
		value = value * 17
		value = int(math.Mod(float64(value), 256))
	}
	return value
}

// too low: 293687
