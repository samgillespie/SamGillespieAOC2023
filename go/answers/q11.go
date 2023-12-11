package answers

func Day11() []interface{} {
	data := ReadInputAsStr(11)
	galaxies, empty_rows, empty_columns := parseGalaxies(data)
	return []interface{}{
		q11part1(galaxies, empty_rows, empty_columns),
		q11part2(galaxies, empty_rows, empty_columns),
	}
}

func parseGalaxies(data []string) ([]Vector, []int, []int) {
	// Returns []Vector of galaxies, empty rows and empty columns
	galaxies := []Vector{}
	filled_rows := []int{}
	filled_columns := []int{}
	max_y := len(data)
	max_x := len(data[0])
	for y, row := range data {
		for x, cell := range row {
			if cell != '#' {
				continue
			}
			galaxies = append(galaxies, Vector{x: x, y: y})
			filled_rows = append(filled_rows, y)
			filled_columns = append(filled_columns, x)
		}
	}
	empty_rows := []int{}
	empty_columns := []int{}
	for x := 0; x < max_x; x++ {
		if IntInSlice(x, filled_rows) {
			continue
		}
		empty_columns = append(empty_columns, x)
	}

	for y := 0; y < max_y; y++ {
		if IntInSlice(y, filled_columns) {
			continue
		}
		empty_rows = append(empty_rows, y)
	}
	return galaxies, empty_rows, empty_columns
}

func q11part1(galaxies []Vector, empty_rows []int, empty_columns []int) int {

	total_distance := 0
	for i, a := range galaxies {
		for j, b := range galaxies {
			if j <= i {
				continue
			}
			distance := GalaxyDistance(a, b, empty_rows, empty_columns, 2)
			total_distance += distance
		}
	}
	return total_distance
}

func GalaxyDistance(a Vector, b Vector, empty_rows []int, empty_columns []int, size_if_empty int) int {
	min_x := min(a.x, b.x)
	max_x := max(a.x, b.x)
	x_dist := 0

	min_y := min(a.y, b.y)
	max_y := max(a.y, b.y)
	y_dist := 0
	for x := min_x; x < max_x; x++ {
		if IntInSlice(x, empty_rows) {
			x_dist += size_if_empty
		} else {
			x_dist += 1
		}
	}

	for y := min_y; y < max_y; y++ {
		if IntInSlice(y, empty_columns) {
			y_dist += size_if_empty
		} else {
			y_dist += 1
		}
	}

	return x_dist + y_dist
}

func q11part2(galaxies []Vector, empty_rows []int, empty_columns []int) int {
	total_distance := 0
	for i, a := range galaxies {
		for j, b := range galaxies {
			if j <= i {
				continue
			}
			distance := GalaxyDistance(a, b, empty_rows, empty_columns, 1000000)
			total_distance += distance
		}
	}
	return total_distance
}

//wrong: 9189792
