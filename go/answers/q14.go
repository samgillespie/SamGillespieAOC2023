package answers

import (
	"fmt"
	"math"
)

func Day14() []interface{} {
	data := ReadInputAsStr(14)
	rock_map := convert_to_rune_map(data)
	return []interface{}{q14part1(copy_to_new_map(rock_map)), q14part2(copy_to_new_map(rock_map))}
}

func copy_to_new_map(rock_map [][]rune) [][]rune {
	new_slice := make([][]rune, len(rock_map))
	for i, row := range rock_map {
		new_row := make([]rune, len(row))
		for j, cell := range row {
			new_row[j] = cell
		}
		new_slice[i] = new_row
	}
	return new_slice
}

func convert_to_rune_map(string_slice []string) [][]rune {
	new_slice := make([][]rune, len(string_slice))
	for i, j := range string_slice {
		new_slice[i] = []rune(j)
	}
	return new_slice
}

func roll_north(rock_map [][]rune) [][]rune {
	rock_locations := []Vector{}
	for y, row := range rock_map {
		for x, cell := range row {
			if cell == 'O' {
				rock_locations = append(rock_locations, Vector{x, y})
			}
		}
	}

	// We can resolve in order
	for _, rock := range rock_locations {
		for y := rock.y; y > 0; y-- {
			if rock_map[y-1][rock.x] == '.' {
				rock_map[y-1][rock.x] = 'O'
				rock_map[y][rock.x] = '.'
			} else {
				break
			}
		}
	}
	return rock_map
}

func roll_east(rock_map [][]rune) [][]rune {
	rock_locations := []Vector{}
	for x := len(rock_map[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(rock_map); y++ {
			cell := rock_map[y][x]
			if cell == 'O' {
				rock_locations = append(rock_locations, Vector{x, y})
			}
		}
	}

	// We can resolve in order
	for _, rock := range rock_locations {
		for x := rock.x; x < len(rock_map[0])-1; x++ {
			if rock_map[rock.y][x+1] == '.' {
				rock_map[rock.y][x+1] = 'O'
				rock_map[rock.y][x] = '.'
			} else {
				break
			}
		}
	}
	return rock_map
}

func roll_west(rock_map [][]rune) [][]rune {
	rock_locations := []Vector{}
	for x := 0; x < len(rock_map[0]); x++ {
		for y := 0; y < len(rock_map); y++ {
			cell := rock_map[y][x]
			if cell == 'O' {
				rock_locations = append(rock_locations, Vector{x, y})
			}
		}
	}

	// We can resolve in order
	for _, rock := range rock_locations {
		for x := rock.x; x > 0; x-- {
			if rock_map[rock.y][x-1] == '.' {
				rock_map[rock.y][x-1] = 'O'
				rock_map[rock.y][x] = '.'
			} else {
				break
			}
		}
	}

	return rock_map
}

func roll_south(rock_map [][]rune) [][]rune {
	rock_locations := []Vector{}
	for y := len(rock_map) - 1; y >= 0; y-- {
		for x := 0; x < len(rock_map[0]); x++ {
			cell := rock_map[y][x]
			if cell == 'O' {
				rock_locations = append(rock_locations, Vector{x, y})
			}
		}
	}
	// We can resolve in order
	for _, rock := range rock_locations {
		for y := rock.y; y < len(rock_map)-1; y++ {
			if rock_map[y+1][rock.x] == '.' {
				rock_map[y+1][rock.x] = 'O'
				rock_map[y][rock.x] = '.'
			} else {
				break
			}
		}
	}

	return rock_map
}

func calculate_score(rock_map [][]rune) int {
	total_score := 0

	max_row := len(rock_map)
	for y, row := range rock_map {
		score := max_row - y
		for _, cell := range row {
			if cell == 'O' {
				total_score += score
			}
		}
	}
	return total_score
}
func q14part1(rock_map [][]rune) int {
	rock_map = roll_north(rock_map)
	return calculate_score(rock_map)
}

func print_map(rock_map [][]rune) {
	for _, row := range rock_map {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func rock_cycle(rock_map [][]rune) [][]rune {
	rock_map = roll_north(rock_map)
	rock_map = roll_west(rock_map)
	rock_map = roll_south(rock_map)
	rock_map = roll_east(rock_map)
	return rock_map
}

func check_for_repeated_pattern_at_chunk_size(sequence []int, chunk_size int) int {
	var previous_chunk []int
	var current_chunk []int
	pattern_observed := 0
	for chunk_num := 0; chunk_num < len(sequence)/chunk_size; chunk_num++ {

		first := chunk_num * chunk_size
		second := (chunk_num + 1) * chunk_size
		if chunk_num == 0 {
			previous_chunk = sequence[first:second]
			continue
		}
		current_chunk = sequence[first:second]
		if IntSliceEqual(previous_chunk, current_chunk) {
			pattern_observed += 1
		}

		// Make sure the pattern occurs 3 times, just to make sure
		if pattern_observed == 3 {
			modulus := math.Mod(1000000000, float64(chunk_size))
			return current_chunk[int(modulus)-1]
		}
		previous_chunk = current_chunk
	}
	return -1
}

func check_for_repeated_pattern(sequence []int) int {
	for interval_size := 2; interval_size < len(sequence)/4; interval_size++ {
		result := check_for_repeated_pattern_at_chunk_size(sequence, interval_size)
		if result != -1 {
			return result
		}
	}
	return -1
}

func q14part2(rock_map [][]rune) int {
	scores := []int{}
	for i := 0; i < 1000; i++ {
		rock_map = rock_cycle(rock_map)
		scores = append(scores, calculate_score(rock_map))
		result := check_for_repeated_pattern(scores)
		if result != -1 {
			return result
		}
	}

	return 0
}
