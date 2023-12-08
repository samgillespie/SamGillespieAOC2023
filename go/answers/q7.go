package answers

import (
	"sort"
	"strconv"
	"strings"
)

var CARD_RANKING_PART_A map[byte]int = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var CARD_RANKING_PART_B map[byte]int = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func Day7() []interface{} {
	data := ReadInputAsStr(7)
	score_map := map[string]int{}
	hands := []string{}
	for _, row := range data {
		split := strings.Split(row, " ")
		score_map[split[0]], _ = strconv.Atoi(split[1])
		hands = append(hands, split[0])
	}
	return []interface{}{q7part1(score_map, hands), q7part2(score_map, hands)}
}

func is_x_of_a_kind(hand string, n int, jokers_wild bool) bool {
	counter := map[rune]int{}
	jokers := 0
	for _, card := range hand {
		if jokers_wild && card == 'J' {
			jokers += 1
		}
		value, exists := counter[card]
		if exists {
			counter[card] = value + 1
		} else {
			counter[card] = 1
		}
	}

	for card, num := range counter {
		if jokers_wild && card != 'J' && num+jokers == n {
			return true
		}
		if num == n {
			return true
		}
	}
	return false
}

func is_two_pair(hand string, jokers_wild bool) bool {
	counter := map[rune]int{}
	jokers := 0
	for _, card := range hand {
		if jokers_wild && card == 'J' {
			jokers += 1
			continue
		}
		value, exists := counter[card]
		if exists {
			counter[card] = value + 1
		} else {
			counter[card] = 1
		}
	}

	pairs := 0
	for _, num := range counter {
		if num == 2 {
			pairs += 1
		}
		if num == 1 && jokers > 0 {
			jokers--
			pairs += 1
		}
		if pairs == 2 {
			return true
		}
	}
	return false
}

func is_full_house(hand string, jokers_wild bool) bool {
	if jokers_wild == false {
		return is_x_of_a_kind(hand, 3, false) && is_x_of_a_kind(hand, 2, false)
	}

	// Count jokers
	jokers := 0
	for _, card := range hand {
		if card == 'J' {
			jokers += 1
		}
	}
	if jokers >= 2 {
		return false
	}
	if is_two_pair(hand, false) && jokers == 1 {
		return true
	}
	return is_x_of_a_kind(hand, 3, false) && is_x_of_a_kind(hand, 2, false)
}

func rank_hand(hand string, jokers_wild bool) int {
	if is_x_of_a_kind(hand, 5, jokers_wild) {
		return 6
	}
	if is_x_of_a_kind(hand, 4, jokers_wild) {
		return 5
	}
	if is_full_house(hand, jokers_wild) {
		return 4
	}
	if is_x_of_a_kind(hand, 3, jokers_wild) {
		return 3
	}
	if is_two_pair(hand, jokers_wild) {
		return 2
	}
	if is_x_of_a_kind(hand, 2, jokers_wild) {
		return 1
	}
	return 0
}

func sort_hand_by_card_value(hand1 string, hand2 string, jokers_wild bool) bool {
	for i := 0; i < 5; i++ {
		var card1, card2 int
		if jokers_wild == false {
			card1 = CARD_RANKING_PART_A[hand1[i]]
			card2 = CARD_RANKING_PART_A[hand2[i]]
		} else {
			card1 = CARD_RANKING_PART_B[hand1[i]]
			card2 = CARD_RANKING_PART_B[hand2[i]]
		}
		if card1 != card2 {
			return card1 < card2
		}
	}
	return false
}

func q7part1(score_map map[string]int, hands []string) int {
	jokers_wild := false
	sort.SliceStable(hands, func(i, j int) bool {
		hand1 := rank_hand(hands[i], jokers_wild)
		hand2 := rank_hand(hands[j], jokers_wild)
		if hand1 == hand2 {
			return sort_hand_by_card_value(hands[i], hands[j], jokers_wild)
		}
		return hand1 < hand2
	})

	total := 0
	for rank, hand := range hands {
		score := score_map[hand]
		total += (rank + 1) * score
	}

	return total
}

func q7part2(score_map map[string]int, hands []string) int {
	jokers_wild := true
	sort.SliceStable(hands, func(i, j int) bool {
		hand1 := rank_hand(hands[i], jokers_wild)
		hand2 := rank_hand(hands[j], jokers_wild)
		if hand1 == hand2 {
			return sort_hand_by_card_value(hands[i], hands[j], jokers_wild)
		}
		return hand1 < hand2
	})

	total := 0
	for rank, hand := range hands {
		score := score_map[hand]
		total += (rank + 1) * score
	}

	return total
}
