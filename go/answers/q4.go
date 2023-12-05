package answers

import (
	"math"
	"strconv"
	"strings"
)

func Day4() []interface{} {
	data := ReadInputAsStr(4)
	scratchcards := make([]Scratchcard, len(data))
	for i, card := range data {
		scratchcards[i] = parse_scratchcard(card)
	}
	return []interface{}{
		q4part1(scratchcards),
		q4part2(scratchcards),
	}
}

type Scratchcard struct {
	number           int
	winning_numbers  []int
	selected_numbers []int
}

func (s Scratchcard) Matches() int {
	matches := 0
	for _, num := range s.selected_numbers {
		if IntInSlice(num, s.winning_numbers) {
			matches += 1
		}
	}
	return matches
}

func (s Scratchcard) Score() int {
	matches := float64(s.Matches())
	return int(math.Pow(2.0, matches-1))
}

func parse_scratchcard(card string) Scratchcard {
	card_split := strings.Split(card, ":")
	card_number, _ := strconv.Atoi(strings.Split(card, " ")[1])

	numbers_split := strings.Split(card_split[1], "|")
	winning_numbers := toListOfInts(strings.Split(numbers_split[0], " "))
	selected_numbers := toListOfInts(strings.Split(numbers_split[1], " "))
	return Scratchcard{
		number:           card_number,
		winning_numbers:  winning_numbers,
		selected_numbers: selected_numbers,
	}
}

func q4part1(cards []Scratchcard) int {
	score := 0
	for _, card := range cards {
		score += card.Score()
	}
	return score
}

func q4part2(cards []Scratchcard) int {
	totals := make([]int, len(cards))

	// Set to 1
	for i := range totals {
		totals[i] = 1
	}

	for i := range totals {
		card := totals[i]
		matches := cards[i].Matches()
		for j := 0; j < matches; j++ {
			if i+j+1 > len(cards) {
				break
			}
			totals[i+j+1] += card
		}
	}
	sum := 0
	for _, i := range totals {
		sum += i
	}
	return sum
}
