package answers

import (
	"fmt"
)

func Day16() []interface{} {
	data := ReadInputAsStr(16)
	return []interface{}{q16part1(data), q16part2(data)}
}

type Beam struct {
	x   int
	y   int
	dir rune // 'N', 'E', 'S', 'W'
}

func (b Beam) hash() int {
	// Terrible hashing algo
	// But should be unique enough for this problem
	hash := b.y*10000 + b.x*10000000 + int(b.dir)
	return hash
}

func print_visited_cells(visited_cells [][]rune, locations []Beam) {
	for _, location := range locations {
		var icon rune
		if location.dir == 'N' {
			icon = '^'
		}
		if location.dir == 'S' {
			icon = 'v'
		}
		if location.dir == 'E' {
			icon = '>'
		}
		if location.dir == 'W' {
			icon = '<'
		}
		visited_cells[location.y][location.x] = icon
	}
	for _, i := range visited_cells {
		fmt.Println(string(i))
	}
	fmt.Println()
}

func count_visted_cells(visited_cells [][]rune) int {
	counter := 0
	for _, i := range visited_cells {
		for _, j := range i {
			if j == '#' {
				counter++
			}
		}
	}
	return counter
}

func run_beam_simulation(data []string, beams []Beam) int {
	x_len := len(data[0])
	y_len := len(data)
	visited_cells := make([][]rune, y_len)
	visited_beams := map[int]bool{}
	for y := 0; y < y_len; y++ {
		visited_cells[y] = make([]rune, x_len)
		for x := 0; x < x_len; x++ {
			visited_cells[y][x] = '.'
		}
	}

	for len(beams) > 0 {
		next_iteration := []Beam{}
		for _, beam := range beams {
			if beam.x >= 0 && beam.x < x_len && beam.y >= 0 && beam.y < y_len {
				visited_cells[beam.y][beam.x] = '#'
			}

			// If we have visited this path before, bail out
			_, has_visited := visited_beams[beam.hash()]
			if has_visited == true {
				continue
			}
			visited_beams[beam.hash()] = true

			var target_x int
			var target_y int
			if beam.dir == 'E' {
				if beam.x+1 >= x_len {
					continue
				}
				target_x = beam.x + 1
				target_y = beam.y
			} else if beam.dir == 'N' {
				if beam.y-1 < 0 {
					continue
				}
				target_x = beam.x
				target_y = beam.y - 1
			} else if beam.dir == 'W' {
				if beam.x-1 < 0 {
					continue
				}
				target_x = beam.x - 1
				target_y = beam.y

			} else if beam.dir == 'S' {
				if beam.y+1 >= y_len {
					continue
				}
				target_x = beam.x
				target_y = beam.y + 1
			} else {
				panic(fmt.Sprintf("Invalid direction: %s", string(beam.dir)))
			}
			cell := data[target_y][target_x]
			visited_cells[target_y][target_x] = '#'

			if cell == '/' {
				var new_dir rune
				switch {
				case beam.dir == 'E':
					new_dir = 'N'
				case beam.dir == 'N':
					new_dir = 'E'
				case beam.dir == 'S':
					new_dir = 'W'
				case beam.dir == 'W':
					new_dir = 'S'
				}
				next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: new_dir})
			} else if cell == '\\' {
				var new_dir rune
				switch {
				case beam.dir == 'E':
					new_dir = 'S'
				case beam.dir == 'N':
					new_dir = 'W'
				case beam.dir == 'S':
					new_dir = 'E'
				case beam.dir == 'W':
					new_dir = 'N'
				}
				next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: new_dir})
			} else if cell == '-' {
				if beam.dir == 'N' || beam.dir == 'S' {
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: 'W'})
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: 'E'})
				} else {
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: beam.dir})
				}
			} else if cell == '|' {
				if beam.dir == 'E' || beam.dir == 'W' {
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: 'N'})
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: 'S'})
				} else {
					next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: beam.dir})
				}
			} else {
				next_iteration = append(next_iteration, Beam{x: target_x, y: target_y, dir: beam.dir})
			}
		}
		beams = next_iteration
	}
	return count_visted_cells(visited_cells)
}

func q16part1(data []string) int {
	beams := []Beam{{x: -1, y: 0, dir: 'E'}}
	return run_beam_simulation(data, beams)
}

func q16part2(data []string) int {
	maximum := 0

	x_len := len(data[0])
	y_len := len(data)
	for x := 0; x < x_len; x++ {
		beams := []Beam{{x: x, y: -1, dir: 'S'}}
		result := run_beam_simulation(data, beams)
		if result > maximum {
			maximum = result
		}

		beams = []Beam{{x: x, y: y_len, dir: 'N'}}
		result = run_beam_simulation(data, beams)
		if result > maximum {
			maximum = result
		}
	}

	for y := 0; y < y_len; y++ {
		beams := []Beam{{x: -1, y: y, dir: 'E'}}
		result := run_beam_simulation(data, beams)
		if result > maximum {
			maximum = result
		}

		beams = []Beam{{x: x_len, y: y, dir: 'W'}}
		result = run_beam_simulation(data, beams)
		if result > maximum {
			maximum = result
		}
	}
	return maximum
}
