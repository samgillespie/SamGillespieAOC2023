package answers

import "fmt"

func Day10() []interface{} {
	data := ReadInputAsStr(10)
	pipes, start := parse_pipes(data)

	return []interface{}{q10part1(pipes, start), q10part2(pipes, start)}
}

type Pipe struct {
	x            int
	y            int
	id           int
	value        rune
	exits        []*Pipe
	is_main_loop bool
}

func parse_pipes(input []string) ([]*Pipe, *Pipe) {
	x_len := len(input[0])
	y_len := len(input)
	pipes := make([]*Pipe, len(input)*len(input[0]))
	var start *Pipe
	var startIndex int
	for y, row := range input {
		for x, rune_elem := range row {
			index := y*x_len + x
			pipes[index] = &Pipe{
				x:     x,
				y:     y,
				id:    index,
				value: rune_elem,
			}
			if rune_elem == 'S' {
				start = pipes[index]
			}
		}
	}
	// Now find exits

	for i, elem := range pipes {
		// North
		if elem.value == '|' || elem.value == 'L' || elem.value == 'J' {
			ynew := elem.y - 1
			if ynew >= 0 {
				index := ynew*x_len + elem.x
				elem.exits = append(elem.exits, pipes[index])
			}
		}
		// East
		if elem.value == '-' || elem.value == 'L' || elem.value == 'F' {
			xnew := elem.x + 1
			if xnew < x_len {
				index := elem.y*x_len + xnew
				elem.exits = append(elem.exits, pipes[index])
			}
		}
		// South
		if elem.value == '|' || elem.value == '7' || elem.value == 'F' {
			ynew := elem.y + 1
			if ynew < y_len {
				index := ynew*x_len + elem.x
				elem.exits = append(elem.exits, pipes[index])
			}
		}
		// West
		if elem.value == '-' || elem.value == '7' || elem.value == 'J' {
			xnew := elem.x - 1
			if xnew >= 0 {
				index := elem.y*x_len + xnew
				elem.exits = append(elem.exits, pipes[index])
			}
		}

		for _, exit := range elem.exits {
			if exit.value == 'S' {
				start.exits = append(start.exits, elem)
				pipes[startIndex] = start
			}
		}
		pipes[i] = elem
	}
	return pipes, start
}

func q10part1(pipes []*Pipe, start *Pipe) int {
	steps := 0
	previous_step := start
	walker := start.exits[0]
	for {
		if walker.exits[0] == previous_step {
			previous_step = walker
			walker = walker.exits[1]
		} else {
			previous_step = walker
			walker = walker.exits[0]
		}
		steps += 1
		if walker.value == 'S' {
			return steps/2 + 1
		}
	}
}

func q10part2(pipes []*Pipe, start *Pipe) int {
	steps := 0
	previous_step := start
	walker := start.exits[0]
	for {
		if walker.exits[0] == previous_step {
			previous_step = walker
			walker = walker.exits[1]
		} else {
			previous_step = walker
			walker = walker.exits[0]
		}
		steps += 1
		walker.is_main_loop = true
		if walker.value == 'S' {
			break
		}
	}
	PrintLoops(pipes)
	return 0
}

func PrintLoops(pipes []*Pipe) {
	row := ""
	for _, pipe := range pipes {
		if pipe.x == 0 {
			fmt.Println(row)
			row = ""
		}
		if pipe.is_main_loop {
			row += "X"
		} else {
			row += "."
		}
	}
}
