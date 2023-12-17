package answers

import "fmt"

func Day17() []interface{} {
	data := ReadInputAsStr(17)
	heats := make([][]int, len(data))
	for y, row := range data {
		heat_row := make([]int, len(row))
		for x, elem := range row {
			heat_row[x] = int(elem) - 48
		}
		heats[y] = heat_row
	}
	return []interface{}{q17part1(heats), q17part2(heats)}
}

type Crucible struct {
	x           int
	y           int
	heat        int
	x_direction int
	y_direction int
}

func min_heat(crucibles []Crucible) int {
	min_heat := 9999999
	for _, elem := range crucibles {
		if elem.heat < min_heat {
			min_heat = elem.heat
		}
	}
	return min_heat
}

func print_heats(heats [][][]Crucible) {
	for _, row := range heats {
		row_value := []int{}
		for _, cell := range row {
			minimum := min_heat(cell)
			if minimum == 9999999 {
				minimum = -1
			}
			row_value = append(row_value, minimum)
		}

		fmt.Println(row_value)
	}
	fmt.Println()
}

func vertex_is_valid(vertex Crucible, previous_vertexes [][][]Crucible) (int, bool) {
	// Returns the position of the matching vertex
	for idx, prev_vertex := range previous_vertexes[vertex.y][vertex.x] {
		if vertex.x_direction == prev_vertex.x_direction && vertex.y_direction == prev_vertex.y_direction {
			return idx, vertex.heat < prev_vertex.heat
		}
	}
	return -1, true
}

func q17part1(data [][]int) int {
	vertices := []Crucible{{x: 0, y: 0}}
	y_len := len(data)
	x_len := len(data[0])

	// Prepare min_heats
	previous_vertices := make([][][]Crucible, y_len)
	for y := 0; y < y_len; y++ {
		previous_vertex_list := make([][]Crucible, x_len)
		for x := 0; x < x_len; x++ {
			previous_vertex_list[x] = []Crucible{}
		}
		previous_vertices[y] = previous_vertex_list
	}
	previous_vertices[0][0] = append(previous_vertices[0][0], Crucible{x: 0, y: 0, heat: 0})

	for {
		if len(vertices) == 0 {
			break
		}
		new_vertices := []Crucible{}
		for _, vertex := range vertices {
			// Left

			directions := [][]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			}
			for _, direction := range directions {

				boundary_left := direction[0] == -1 && vertex.x > 0 && vertex.x_direction > -3 && vertex.x_direction <= 0
				boundary_right := direction[0] == 1 && vertex.x < x_len-1 && vertex.x_direction < 3 && vertex.x_direction >= 0
				boundary_up := direction[1] == -1 && vertex.y > 0 && vertex.y_direction > -3 && vertex.y_direction <= 0
				boundary_down := direction[1] == 1 && vertex.y < y_len-1 && vertex.y_direction < 3 && vertex.y_direction >= 0

				if boundary_left || boundary_down || boundary_right || boundary_up {
					new_x := vertex.x + direction[0]
					new_y := vertex.y + direction[1]
					new_vertex := Crucible{
						x:    new_x,
						y:    new_y,
						heat: vertex.heat + data[new_y][new_x],
					}

					if direction[0] == -1 {
						if vertex.x_direction < 0 {
							new_vertex.x_direction = vertex.x_direction + direction[0]
						} else {
							new_vertex.x_direction = direction[0]
						}
					}

					if direction[0] == 1 {
						if vertex.x_direction > 0 {
							new_vertex.x_direction = vertex.x_direction + direction[0]
						} else {
							new_vertex.x_direction = direction[0]
						}
					}

					if direction[1] == -1 {
						if vertex.y_direction < 0 {
							new_vertex.y_direction = vertex.y_direction + direction[1]
						} else {
							new_vertex.y_direction = direction[1]
						}
					}

					if direction[1] == 1 {
						if vertex.y_direction > 0 {
							new_vertex.y_direction = vertex.y_direction + direction[1]
						} else {
							new_vertex.y_direction = direction[1]
						}
					}

					idx, is_valid := vertex_is_valid(new_vertex, previous_vertices)
					// fmt.Println(idx, is_valid)
					if is_valid {
						if idx != -1 {
							previous_vertices[new_y][new_x][idx] = new_vertex
						} else {
							previous_vertices[new_y][new_x] = append(previous_vertices[new_y][new_x], new_vertex)
						}
						new_vertices = append(new_vertices, new_vertex)
					}
				}
			}
		}
		vertices = new_vertices
		// print_heats(previous_vertices)
	}
	return min_heat(previous_vertices[y_len-1][x_len-1])
}

func q17part2(data [][]int) int {
	vertices := []Crucible{{x: 0, y: 0}}
	y_len := len(data)
	x_len := len(data[0])

	// Prepare min_heats
	previous_vertices := make([][][]Crucible, y_len)
	for y := 0; y < y_len; y++ {
		previous_vertex_list := make([][]Crucible, x_len)
		for x := 0; x < x_len; x++ {
			previous_vertex_list[x] = []Crucible{}
		}
		previous_vertices[y] = previous_vertex_list
	}
	previous_vertices[0][0] = append(previous_vertices[0][0], Crucible{x: 0, y: 0, heat: 0})

	for {
		if len(vertices) == 0 {
			break
		}
		new_vertices := []Crucible{}
		for _, vertex := range vertices {
			// Left

			directions := [][]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			}
			for _, direction := range directions {
				boundary_left := direction[0] == -1 &&
					vertex.x > 0 &&
					vertex.x+abs(vertex.x_direction) >= 4 && // Cannot start going left if I'm at position 3 or less
					vertex.x_direction > -10 && // Maximum in one direction
					(vertex.x_direction < 0 || abs(vertex.y_direction) >= 4) // already moving, or y finished mobing

				boundary_right := direction[0] == 1 &&
					vertex.x-abs(vertex.x_direction) < x_len-4 &&
					vertex.x < x_len-1 &&
					vertex.x_direction < 10 &&
					(vertex.x_direction > 0 || abs(vertex.y_direction) >= 4 || (vertex.x_direction == 0 && vertex.y_direction == 0))

				boundary_up := direction[1] == -1 &&
					vertex.y+abs(vertex.y_direction) >= 4 &&
					vertex.y > 0 &&
					vertex.y_direction > -10 &&
					(vertex.y_direction < 0 || abs(vertex.x_direction) >= 4)

				boundary_down := direction[1] == 1 &&
					vertex.y-abs(vertex.y_direction) < y_len-4 &&
					vertex.y < y_len-1 &&
					vertex.y_direction < 10 &&
					(vertex.y_direction > 0 || abs(vertex.x_direction) >= 4 || (vertex.x_direction == 0 && vertex.y_direction == 0))

				if boundary_left || boundary_down || boundary_right || boundary_up {
					new_x := vertex.x + direction[0]
					new_y := vertex.y + direction[1]
					new_vertex := Crucible{
						x:    new_x,
						y:    new_y,
						heat: vertex.heat + data[new_y][new_x],
					}

					if direction[0] == -1 {
						if vertex.x_direction < 0 {
							new_vertex.x_direction = vertex.x_direction + direction[0]
						} else {
							new_vertex.x_direction = direction[0]
						}
					}

					if direction[0] == 1 {
						if vertex.x_direction > 0 {
							new_vertex.x_direction = vertex.x_direction + direction[0]
						} else {
							new_vertex.x_direction = direction[0]
						}
					}

					if direction[1] == -1 {
						if vertex.y_direction < 0 {
							new_vertex.y_direction = vertex.y_direction + direction[1]
						} else {
							new_vertex.y_direction = direction[1]
						}
					}

					if direction[1] == 1 {
						if vertex.y_direction > 0 {
							new_vertex.y_direction = vertex.y_direction + direction[1]
						} else {
							new_vertex.y_direction = direction[1]
						}
					}

					idx, is_valid := vertex_is_valid(new_vertex, previous_vertices)
					if is_valid {
						if idx != -1 {
							previous_vertices[new_y][new_x][idx] = new_vertex
						} else {
							previous_vertices[new_y][new_x] = append(previous_vertices[new_y][new_x], new_vertex)
						}
						new_vertices = append(new_vertices, new_vertex)
					}
				}
			}
		}
		vertices = new_vertices
	}
	return min_heat(previous_vertices[y_len-1][x_len-1])
}
