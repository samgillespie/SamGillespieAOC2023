package answers

import (
	"strconv"
	"strings"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

func (c Cubes) IsPart1Valid() bool {
	// Returns false if it has more than 12 red, 13 green or 14 blue
	if c.red > 12 {
		return false
	}
	if c.green > 13 {
		return false
	}
	if c.blue > 14 {
		return false
	}
	return true
}

type Game struct {
	number int
	cubes  []Cubes
}

func (g Game) IsPart1Valid() bool {
	// Returns false if it has more than 12 red, 13 green or 14 blue
	for _, cube := range g.cubes {
		if cube.IsPart1Valid() == false {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	min_red := 0
	min_green := 0
	min_blue := 0
	for _, cube := range g.cubes {
		if cube.red > min_red {
			min_red = cube.red
		}
		if cube.green > min_green {
			min_green = cube.green
		}
		if cube.blue > min_blue {
			min_blue = cube.blue
		}
	}
	return min_blue * min_green * min_red
}

func ParseCubes(input string) Cubes {
	// Convert 3 blue, 4 red, 5 green -> Cubes struct
	input_split := strings.Split(input, ",")
	cube := Cubes{}
	for _, elem := range input_split {
		elem_trimmed := strings.Trim(elem, " ")
		entry_split := strings.Split(elem_trimmed, " ")
		number, err := strconv.Atoi(entry_split[0])
		if err != nil {
			panic(err)
		}
		if entry_split[1] == "red" {
			cube.red = number
		} else if entry_split[1] == "green" {
			cube.green = number
		} else if entry_split[1] == "blue" {
			cube.blue = number
		} else {
			panic(elem)
		}
	}
	return cube
}

func ParseGame(input string) Game {
	split_input := strings.Split(input, ":")
	game_number, err := strconv.Atoi(strings.Split(split_input[0], " ")[1])
	if err != nil {
		panic(err)
	}
	game := Game{number: game_number}
	game.cubes = []Cubes{}
	cubes_split := strings.Split(split_input[1], ";")
	for _, cubes := range cubes_split {
		game.cubes = append(game.cubes, ParseCubes(cubes))
	}
	return game
}

func Day2() []interface{} {
	data := ReadInputAsStr(2)
	games := []Game{}
	for _, row := range data {
		games = append(games, ParseGame(row))

	}
	return []interface{}{q2part1(games), q2part2(games)}

}

func q2part1(games []Game) int {
	total := 0
	for _, game := range games {
		if game.IsPart1Valid() {
			total += game.number
		}
	}
	return total
}

func q2part2(games []Game) int {
	power := 0
	for _, game := range games {
		power += game.Power()
	}
	return power
}
