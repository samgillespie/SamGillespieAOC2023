package answers

import (
	"strconv"
	"strings"
)

func Day6() []interface{} {
	data := ReadInputAsStr(6)
	return []interface{}{q6part1(data), q6part2(data)}
}

func q6part1(data []string) int {
	timeArray := toListOfInts(strings.Split(data[0], " ")[1:])
	distanceArray := toListOfInts(strings.Split(data[1], " ")[1:])

	total := 1
	for i := 0; i < len(timeArray); i++ {
		time := timeArray[i]
		threshold := distanceArray[i]
		number := 0
		for j := 1; j < time; j++ {
			distance := j * (time - j)
			if distance > threshold {
				number++
			}
		}
		total = total * number
	}
	return total
}

func q6part2(data []string) int {
	time, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(data[0], " ", ""), ":")[1])
	threshold, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(data[1], " ", ""), ":")[1])
	for i := 0; i < time; i++ {
		distance := i * (time - i)
		if distance > threshold {
			return time - 2*i + 1
		}
	}
	return 0
}
