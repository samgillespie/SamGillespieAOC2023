package main

import (
	"code/aoc2023/answers"
	"flag"
	"fmt"
	"time"
)

var QUESTION int
var runProfile bool

var questionMap = map[int]func() []interface{}{
	1: answers.Day1,
	2: answers.Day2,
	3: answers.Day3,
	// 4:  answers.Day4,
	// 5:  answers.Day5,
	// 6:  answers.Day6,
	// 7:  answers.Day7,
	// 8:  answers.Day8,
	// 9:  answers.Day9,
	// 10: answers.Day10,
	// 11: answers.Day11,
	// 12: answers.Day12,
	// 13: answers.Day13,
	// 14: answers.Day14,
	// 15: answers.Day15,
	// 16: answers.Day16,
	// 17: answers.Day17,
	// 18: answers.Day18,
	// 19: answers.Day19,
	// 20: answers.Day20,
	// 21: answers.Day21,
	// 22: answers.Day22,
	// 23: answers.Day23,
	// 24: answers.Day24,
	// 25: answers.Day25,
}

func main() {
	parseArgs()
	if runProfile == false {
		result := SolveQuestion(false, QUESTION)
		fmt.Printf("Day %d Part 1 Answer : %v\n", QUESTION, result[0])
		fmt.Printf("Day %d Part 2 Answer : %v\n", QUESTION, result[1])
	} else {
		for question := 1; question <= len(questionMap); question++ {
			if question != QUESTION && QUESTION != 0 {
				continue
			}
			runs := make([]time.Duration, 0, 20)
			for i := 0; i < 20; i++ {
				start := time.Now()
				SolveQuestion(true, question)
				runs = append(runs, time.Since(start))
			}
			var min, max, total time.Duration
			for i, runtime := range runs {
				if i == 0 {
					min = runtime
				}
				if runtime < min {
					min = runtime
				}
				if runtime > max {
					max = runtime
				}
				total += runtime
			}
			avg := time.Duration(total.Nanoseconds() / int64(len(runs)))
			fmt.Println("Q", question, "min:", min, "max:", max, "avg:", avg, "total", total)
		}
	}
}

func SolveQuestion(silent bool, question int) []interface{} {
	if question == 0 {
		times := []time.Duration{}
		for i := 1; i <= len(questionMap); i++ {
			start := time.Now()
			questionMap[i]()
			end := time.Since(start)
			times = append(times, end)
			if !silent {
				fmt.Printf("Day %d: Time Taken %s\n", i, end)
			}

		}
		var totalDuration time.Duration
		for _, dur := range times {
			totalDuration += dur
		}
		if !silent {
			fmt.Printf("Total Time Taken: %s\n\n", totalDuration)
		}
		return []interface{}{0, 0}
	} else {
		return questionMap[question]()
	}
}

func parseArgs() {
	flag.IntVar(&QUESTION, "question", 0, "Which question to answer")
	flag.BoolVar(&runProfile, "prof", false, "Whether to run a profile. If enabled runs the solution 1000 times and grabs an average, min and max runtimes")
	flag.Parse()

}

// 544502 Too high
