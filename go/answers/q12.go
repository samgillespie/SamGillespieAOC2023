package answers

import (
	"fmt"
	"strings"
)

func Day12() []interface{} {
	data := ReadInputAsStr(12)
	springs := parse_springs(data)
	return []interface{}{q12part1(springs), q12part2(springs)}
}

type Spring struct {
	str     string
	numbers []int
}

func (s Spring) Unfurl(n int) Spring {
	numbers := make([]int, len(s.numbers)*5)
	for i := 0; i < n; i++ {
		for j := 0; j < len(s.numbers); j++ {
			idx := i*len(s.numbers) + j
			numbers[idx] = s.numbers[j]
		}
	}
	str := ""
	for i := 0; i < n; i++ {
		str += s.str
	}
	return Spring{
		str:     str,
		numbers: numbers,
	}
}

func parse_springs(data []string) []Spring {
	springs := make([]Spring, len(data))
	for i, row := range data {
		row_split := strings.Split(row, " ")

		numbers := strings.Split(row_split[1], ",")
		springs[i] = Spring{
			str:     row_split[0],
			numbers: toListOfInts(numbers),
		}
	}
	return springs
}

func binomial(n, k int) int {
	// How many elements will be in the combinations
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func combinations(n int, k int) [][]int {
	// Code shamelessly stolen from gonum.stat.combin
	combins := binomial(n, k)
	fmt.Println(combins)
	data := make([][]int, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]int, k)
	for i := range data[0] {
		data[0][i] = i
	}
	for i := 1; i < combins; i++ {
		next := make([]int, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

func nextCombination(s []int, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if s[j] == n+j-k {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = s[j] + l - j
		}
		break
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func string_obeys_rule(spring string, number []int) bool {
	spring = strings.ReplaceAll(spring, "?", ".")
	spring_split := strings.Split(spring, ".")
	index := 0
	for _, substring := range spring_split {
		if substring == "" {
			continue
		}
		if len(substring) != number[index] {
			return false
		}
		index++
	}
	return true
}

func find_possible_solutions(spring string, numbers []int) int {
	question_marks := 0
	hashes := 0
	question_mark_locations := []int{}
	for i, rune_elem := range spring {
		if rune_elem == '?' {
			question_marks++
			question_mark_locations = append(question_mark_locations, i)
		}
		if rune_elem == '#' {
			hashes++
		}
	}
	solution_hashes := sumSlice(numbers)
	// fmt.Printf("n=%d, solution_hashes=%d, hashes=%d\n", len(question_mark_locations), solution_hashes, hashes)
	combs := combinations(len(question_mark_locations), solution_hashes-hashes)
	// fmt.Println(combs)
	valid := 0
	for _, comb := range combs {
		new_string := spring
		for _, j := range comb {
			new_string = replaceAtIndex(new_string, '#', question_mark_locations[j])
		}
		// fmt.Println("!!!!!!!!!!!!!!!")
		// fmt.Println(new_string)
		is_valid := string_obeys_rule(new_string, numbers)
		if is_valid {
			valid++
		}

	}
	return valid
}

func test_case() int {
	test := 0
	test += find_possible_solutions("???.###", []int{1, 1, 3})
	test += find_possible_solutions(".??..??...?##.", []int{1, 1, 3})
	test += find_possible_solutions("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6})
	test += find_possible_solutions("????.#...#...", []int{4, 1, 1})
	test += find_possible_solutions("????.######..#####", []int{1, 6, 5})
	test += find_possible_solutions("?###????????", []int{3, 2, 1})
	return test
}

func q12part1(springs []Spring) int {
	return 0
	total := 0
	for _, spring := range springs {
		total += find_possible_solutions(spring.str, spring.numbers)
		//fmt.Println(spring.str, spring.numbers, total)
	}

	return total
}

func find_possible_solutions_using_chunking(spring string, numbers []int) int {
	chunks := strings.Split(spring, ".")
	active_numbers := numbers
	for _, chunk := range chunks {
		chunk_length := len(chunk)

		// pull from active_numbers
	}
	return 0
}

// ?????.??#??.????,  1,3,4

func q12part2(data []Spring) int {
	// a := data[0].Unfurl(5)
	find_possible_solutions("???.###", []int{1})
	return 0
}
