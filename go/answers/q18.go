package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day18() []interface{} {
	data := ReadInputAsStr(18)
	dig_plan := ParseDigPlan(data)
	return []interface{}{q18part1(dig_plan), q18part2(dig_plan)}
}

type DigPlan struct {
	dir      byte
	distance int
	hex_code string
}

func (d DigPlan) parse_hex_code() (rune, int) {
	// Converts the hex_code to a direction and a distance
	dir := ' '
	switch {
	case d.hex_code[5] == '0':
		dir = 'R'
	case d.hex_code[5] == '1':
		dir = 'D'
	case d.hex_code[5] == '2':
		dir = 'L'
	case d.hex_code[5] == '3':
		dir = 'U'
	default:
		panic("Error parsing hex_code")
	}
	decimal_num, err := strconv.ParseInt(d.hex_code[0:5], 16, 64)
	if err != nil {
		panic(err)
	}
	return dir, int(decimal_num)
}

func ParseDigPlan(data []string) []DigPlan {
	result := make([]DigPlan, len(data))
	for x, row := range data {
		split_string := strings.Split(row, " ")
		distance, _ := strconv.Atoi(split_string[1])
		result[x] = DigPlan{
			dir:      split_string[0][0],
			distance: distance,
			hex_code: split_string[2][2:8],
		}
	}
	return result
}

func relative_x(x_pos int, x_min int) int {
	return x_pos - x_min
}

func relative_y(y_pos int, y_min int) int {
	return y_pos - y_min
}

func print_rune_map(rune_map [][]rune) {
	for _, row := range rune_map {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func fill_cell(rune_map *[][]rune, x int, y int) {
	rune_values := (*rune_map)
	if rune_values[y][x] == '.' {
		rune_values[y][x] = 'X'
	}

	if x < len(rune_values[0])-1 {
		if rune_values[y][x+1] == '.' {
			fill_cell(rune_map, x+1, y)
		}
	}
	if y < len(rune_values)-1 {
		if rune_values[y+1][x] == '.' {
			fill_cell(rune_map, x, y+1)
		}
	}

	if x > 0 {
		if rune_values[y][x-1] == '.' {
			fill_cell(rune_map, x-1, y)
		}
	}
	if y > 0 {
		if rune_values[y-1][x] == '.' {
			fill_cell(rune_map, x, y-1)
		}
	}
}

func fill_rune_map(rune_map [][]rune) int {
	fill_cell(&rune_map, 0, 0)
	// print_rune_map(rune_map)
	counter := 0
	for _, row := range rune_map {
		for _, cell := range row {
			if cell != 'X' {
				counter++
			}
		}
	}
	return counter
}

func add_row_at_top(dig_map [][]rune, x_min int, x_max int) [][]rune {
	new_row := make([]rune, x_max-x_min)
	for i := range new_row {
		new_row[i] = '.'
	}
	dig_map = append([][]rune{new_row}, dig_map...)
	return dig_map
}

func add_row_at_bottom(dig_map [][]rune, x_min int, x_max int) [][]rune {
	new_row := make([]rune, x_max-x_min)
	for i := range new_row {
		new_row[i] = '.'
	}
	dig_map = append(dig_map, new_row)
	return dig_map
}

func add_column_at_left(dig_map [][]rune) [][]rune {
	for x := range dig_map {
		dig_map[x] = append([]rune{'.'}, dig_map[x]...)
	}
	return dig_map
}

func add_column_at_right(dig_map [][]rune) [][]rune {
	for x := range dig_map {
		dig_map[x] = append(dig_map[x], '.')
	}
	return dig_map
}

func q18part1(dig_plan []DigPlan) int {
	dig_map := [][]rune{{'#'}}
	var x_min, y_min, x_pos, y_pos int
	x_max := 1
	y_max := 1
	for _, plan := range dig_plan {
		for steps := 0; steps < plan.distance; steps++ {

			if plan.dir == 'D' {
				if y_pos >= y_max-2 {
					dig_map = add_row_at_bottom(dig_map, x_min, x_max)
					y_max++
				}
				y_pos++
				dig_map[relative_y(y_pos, y_min)][relative_x(x_pos, x_min)] = '#'
			} else if plan.dir == 'U' {
				if y_pos == y_min {
					dig_map = add_row_at_top(dig_map, x_min, x_max)
					y_min--
					y_pos--
				} else {
					y_pos--
				}
				dig_map[relative_y(y_pos, y_min)][relative_x(x_pos, x_min)] = '#'
			} else if plan.dir == 'L' {
				if x_pos == x_min {
					dig_map = add_column_at_left(dig_map)
					x_pos--
					x_min--
				} else {
					x_pos--
				}
				dig_map[relative_y(y_pos, y_min)][relative_x(x_pos, x_min)] = '#'
			} else if plan.dir == 'R' {
				if x_pos >= x_max-1 {
					dig_map = add_column_at_right(dig_map)
					x_max++
				}
				x_pos++
				dig_map[relative_y(y_pos, y_min)][relative_x(x_pos, x_min)] = '#'
			}
		}
	}

	// add padding to support a fill algorithm
	dig_map = add_row_at_top(dig_map, x_min, x_max)
	dig_map = add_row_at_bottom(dig_map, x_min, x_max)
	add_column_at_right(dig_map)
	x_max++
	add_column_at_left(dig_map)
	x_min--

	// print_rune_map(dig_map)
	return fill_rune_map(dig_map)
}

func determinate(a Vector, b Vector) int {
	//  | x1, x2 |
	//  | y1, y2 |
	return a.x*b.y - b.x*a.y
}

func q18part2(dig_plan []DigPlan) int {
	vertices := []Vector{{x: 0, y: 0}}
	var x, y int
	circumference := 0
	for _, plan := range dig_plan {
		dir, distance := plan.parse_hex_code()
		switch {
		case dir == 'U':
			y -= distance
		case dir == 'D':
			y += distance
		case dir == 'L':
			x -= distance
		case dir == 'R':
			x += distance
		}
		vertices = append(vertices, Vector{x: x, y: y})
		circumference += distance
	}

	//Shoelace formula
	total := 0
	for i := 0; i < len(vertices)-1; i++ {
		total += determinate(vertices[i], vertices[i+1])
	}
	total += circumference
	return total/2 + 1
}
