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

type ChunkStage struct {
	chunk_cursor   int
	numbers_cursor int
	permutations   int
}

func (c ChunkStage) Print() {
	// fmt.Printf("chunk_cursor: %d, numbers_cursor: %d, permutations: %d\n", c.chunk_cursor, c.numbers_cursor, c.permutations)
}

func (s Spring) Unfurl(n int) Spring {
	numbers := make([]int, len(s.numbers)*n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(s.numbers); j++ {
			idx := i*len(s.numbers) + j
			numbers[idx] = s.numbers[j]
		}
	}
	str := ""
	for i := 0; i < n; i++ {
		str += s.str
		str += "?"
	}
	str = str[0 : len(str)-1]
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

func pull_from_numbers(numbers []int, maximum int) []int {
	// pulls from numbers until the total > maximum
	new_numbers := []int{}
	total := 0
	for _, i := range numbers {
		total += i
		if total <= maximum {
			new_numbers = append(new_numbers, i)
		}
		total += 1
	}
	return new_numbers
}

func chunk_obeys_rule(spring string, number []int) bool {
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

func find_all_valid_combinations(chunk string, numbers []int) int {
	// works out all possible rearrangements of chunks to satify numbers
	question_marks := 0
	hashes := 0
	question_mark_locations := []int{}
	for i, rune_elem := range chunk {
		if rune_elem == '?' {
			question_marks++
			question_mark_locations = append(question_mark_locations, i)
		}
		if rune_elem == '#' {
			hashes++
		}
	}
	// Boundary conditions
	if len(numbers) == 0 && hashes > 0 {
		return 0
	}
	if len(numbers) == 0 && hashes == 0 {
		return 1
	}
	if question_marks == 0 {
		if len(numbers) == 0 || len(numbers) > 1 {
			return 0
		}
		if hashes == numbers[0] {
			return 1
		}
		return 0
	}

	solution_hashes := sumSlice(numbers)
	if solution_hashes < hashes {
		return 0
	}
	combs := combinations(len(question_mark_locations), solution_hashes-hashes)
	valid := 0
	for _, comb := range combs {
		new_chunk := chunk
		for _, j := range comb {
			new_chunk = replaceAtIndex(new_chunk, '#', question_mark_locations[j])
		}

		is_valid := chunk_obeys_rule(new_chunk, numbers)
		if is_valid {
			// fmt.Println("Chunk ", chunk, " solution ", new_chunk, " for ", numbers)
			valid++
		}

	}
	return valid
}

func split_spring(spring string, numbers []int) []string {
	// Optimization, if we see ??###??, and we know the maximum number is 3
	// then we can replace it with ?.###.?

	_, max_interval := maxSlice(numbers)
	largest_contigous := strings.Repeat("#", max_interval)
	spring = strings.ReplaceAll(spring, largest_contigous+"?", largest_contigous+".")
	spring = strings.ReplaceAll(spring, "?"+largest_contigous, "."+largest_contigous)
	// fmt.Println(spring)
	solution := []string{}
	chunks := strings.Split(spring, ".")

	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
		solution = append(solution, chunk)
	}
	return solution
}

func chunk_contains_question_mark(chunk string) bool {
	return string_contains_rune(chunk, '?')
}
func chunk_contains_hash(chunk string) bool {
	return string_contains_rune(chunk, '#')
}

func find_possible_solutions_using_chunking(spring string, numbers []int) int {
	chunks := split_spring(spring, numbers)
	stages := []ChunkStage{
		{
			chunk_cursor:   0,
			numbers_cursor: 0,
			permutations:   1,
		},
	}
	//// fmt.Println("Calculating chunks ", chunks, " using numbers ", numbers)
	solution := 0
	iter := 0
	for len(stages) > 0 {
		next_step := []ChunkStage{}
		// fmt.Println("!!!!!!!!! ROUND ", iter)
		// fmt.Println(stages)
		for _, stage := range stages {

			/////
			//  Finalize path
			/////
			// IF we have processed all the numbers and all the chunks, finish up
			if stage.numbers_cursor >= len(numbers) && stage.chunk_cursor >= len(chunks) {
				solution += stage.permutations
				// fmt.Println("Finalizing ", stage, "solution currently ", solution)
				continue
			}

			// If we have finished all the chunks, and not finished all the numbers
			// This path is invalid
			if stage.chunk_cursor >= len(chunks) {
				// fmt.Println("Abandoning chunk because not all numbers are satisfied", stage)
				continue
			}

			// If we have any unprocessed numbers, check if the remaining chunks are all question marks
			if stage.numbers_cursor >= len(numbers) {
				// If there's any hashes left in unprocessed chunks,
				// Then this path is invalid
				path_valid := true
				for i := stage.chunk_cursor; i < len(chunks); i++ {
					// fmt.Printf("PATH CHECK FOR %s\n", chunks[i])
					if chunk_contains_hash(chunks[i]) {
						path_valid = false
						break
					}
				}
				if path_valid == false {
					// fmt.Println("Abandoning chunk because all numbers are satified, but unifhsed hashes exist", stage)
					continue
				}
				solution += stage.permutations
				// fmt.Println("Finalizing ", stage, "solution currently ", solution)
				continue
			}

			//////
			// Walk the path
			//////
			chunk := chunks[stage.chunk_cursor]

			// We only want to consider numbers that can fit within the chunk
			number_subset := pull_from_numbers(numbers[stage.numbers_cursor:], len(chunk))
			// fmt.Println("Considering combinations for ", chunk, number_subset)
			for i := 0; i < len(number_subset); i++ {

				// fmt.Println("Considering combinations for ", chunk, number_subset[i:i+1])
				combs := find_all_valid_combinations(chunk, number_subset[0:i+1])

				if combs == 0 {
					// Impossible to achieve
					continue
				}
				next_stage := ChunkStage{
					chunk_cursor:   stage.chunk_cursor + 1,
					numbers_cursor: stage.numbers_cursor + 1 + i,
					permutations:   combs * stage.permutations,
				}
				// fmt.Println("Adding chunk due to  next increment", next_stage)
				next_step = append(next_step, next_stage)
			}

			// Also add skipping the chunk, if there's the option to.
			// Cannot skip chunk if it contains a hash
			if chunk_contains_question_mark(chunk) && !chunk_contains_hash(chunk) {
				next_stage := ChunkStage{
					chunk_cursor:   stage.chunk_cursor + 1,
					numbers_cursor: stage.numbers_cursor,
					permutations:   stage.permutations,
				}
				// fmt.Println("Adding chunk skipping current", next_stage)
				next_step = append(next_step, next_stage)
			}
		}
		// fmt.Println(next_step)
		stages = next_step
		iter++
		if iter > 1000 {
			panic("infinite loop")
		}
	}
	// fmt.Println("SOLUTION FOR ", chunks, " using numbers ", numbers, " is ", solution)
	return solution
}

func brute_force(spring string, numbers []int) int {
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
	combs := combinations(len(question_mark_locations), solution_hashes-hashes)
	valid := 0
	for _, comb := range combs {
		new_string := spring
		for _, j := range comb {
			new_string = replaceAtIndex(new_string, '#', question_mark_locations[j])
		}
		is_valid := chunk_obeys_rule(new_string, numbers)
		if is_valid {
			fmt.Println(new_string)
			valid++
		}

	}
	return valid
}

func q12part1(springs []Spring) int {
	solutions := 0
	for _, spring := range springs {
		solution := find_possible_solutions_using_chunking(spring.str, spring.numbers)
		// fmt.Println(spring.str, solution)
		solutions += solution
	}
	return solutions
}

func q12part2(springs []Spring) int {
	solutions := 0

	for _, spring := range springs {
		// We can take advantage of an interesting property
		// If the base has exactly one solution, than unfurl(x) will also have 1 solution
		solution := find_possible_solutions_using_chunking("?"+spring.str, spring.numbers)
		if solution == 1 {
			solutions += 1
			continue
		}

		new_spring := spring.Unfurl(2)
		brute_force(new_spring.str, new_spring.numbers)
		// fmt.Println(new_spring.str, split_spring(new_spring.str, new_spring.numbers))
		solution = find_possible_solutions_using_chunking(new_spring.str, new_spring.numbers)
		fmt.Println(new_spring.str, new_spring.numbers, solution)
		if solution == 0 {
			// fmt.Println("Invalid solution for spring %d", spring.str)
			return 0
		}
		solutions += solution

	}
	return solutions
}
