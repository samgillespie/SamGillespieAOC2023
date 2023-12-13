package answers

func Day13() []interface{} {
	data := ReadInputAsStr(13)
	lava_maps := parseLava(data)
	return []interface{}{q13part1(lava_maps), q13part2(lava_maps)}
}

func parseLava(data []string) [][]string {
	result := [][]string{}
	current_map := []string{}
	for _, row := range data {
		if row == "" {
			result = append(result, current_map)
			current_map = []string{}
		} else {
			current_map = append(current_map, row)
		}
	}
	result = append(result, current_map)
	return result
}

func rotate_lava_map(lava_map []string) []string {
	rotated := make([]string, len(lava_map[0]))
	for i := 0; i < len(lava_map[0]); i++ {
		column := make([]byte, len(lava_map))
		for j := 0; j < len(lava_map); j++ {
			column[j] = lava_map[j][i]
		}
		rotated[i] = string(column)
	}
	return rotated
}

func find_reflection_row(lava_map []string) int {
	for i := 0; i < len(lava_map)-1; i++ {
		symmetrical := true
		for j := 0; j <= len(lava_map); j++ {
			// Check for out of bounds
			if i-j < 0 || i+j+1 >= len(lava_map) {
				break
			}
			if lava_map[i-j] != lava_map[i+j+1] {
				symmetrical = false
				break
			}
		}
		if symmetrical == true {
			return i + 1
		}
	}
	return -1
}

func hamming_distance(a string, b string) int {
	// Returns the number of different characters
	// Assumes len(a) == len(b)
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

func find_reflection_row_with_smudge(lava_map []string) int {
	for i := 0; i < len(lava_map)-1; i++ {
		symmetrical := true
		smudge_applied := false
		for j := 0; j <= len(lava_map); j++ {
			// Check for out of bounds
			if i-j < 0 || i+j+1 >= len(lava_map) {
				break
			}

			distance := hamming_distance(lava_map[i-j], lava_map[i+j+1])
			if distance >= 2 {
				symmetrical = false
				break
			}
			if distance == 1 && smudge_applied {
				symmetrical = false
				break
			}
			if distance == 1 {
				smudge_applied = true
			}
		}

		if smudge_applied && symmetrical {
			return i + 1
		}
	}
	return -1
}

func find_reflection(lava_map []string) []int {
	reflection_row := find_reflection_row(lava_map)
	if reflection_row != -1 {
		return []int{reflection_row, 0}
	}
	rlava_map := rotate_lava_map(lava_map)
	reflection_column := find_reflection_row(rlava_map)
	return []int{0, reflection_column}
}

func find_reflection_with_smudge(lava_map []string) []int {
	rlava_map := rotate_lava_map(lava_map)
	reflection_column := find_reflection_row_with_smudge(rlava_map)
	if reflection_column != -1 {
		return []int{0, reflection_column}
	}

	reflection_row := find_reflection_row_with_smudge(lava_map)
	return []int{reflection_row, 0}
}

func q13part1(lava_maps [][]string) int {
	solution := 0
	for _, lava_map := range lava_maps {
		reflection := find_reflection(lava_map)
		solution += reflection[0]*100 + reflection[1]
	}
	return solution
}

func q13part2(lava_maps [][]string) int {
	solution := 0
	for _, lava_map := range lava_maps {
		reflection := find_reflection_with_smudge(lava_map)
		solution += reflection[0]*100 + reflection[1]
	}
	return solution
}
