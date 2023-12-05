package answers

import (
	"fmt"
	"sort"
	"strings"
)

func Day5() []interface{} {
	data := ReadInputAsStr(5)
	seedInput := parse_seeds(data)
	return []interface{}{q5part1(seedInput), q5part2(seedInput)}
}

type SeedInput struct {
	seeds                   []int
	seed_to_soil            [][]int
	soil_to_fertilizer      [][]int
	fertilizer_to_water     [][]int
	water_to_light          [][]int
	light_to_temperature    [][]int
	temperature_to_humidity [][]int
	humidity_to_location    [][]int
}

func seedMapLookup(currentIndex int, targetMap []int) int {
	destinationStart := targetMap[0]
	sourceStart := targetMap[1]
	rangeOf := targetMap[2]
	distance_from_source_start := currentIndex - sourceStart
	if distance_from_source_start < rangeOf && distance_from_source_start >= 0 {
		return destinationStart + distance_from_source_start
	}
	return -1
}

func (s SeedInput) GetLocation(index int) int {
	walker := s.seeds[index]

	for _, sts := range s.seed_to_soil {
		target_index := seedMapLookup(walker, sts)
		if target_index != -1 {
			walker = target_index
			break
		}
	}
	// Now we're a soil
	for _, stf := range s.soil_to_fertilizer {
		target_index := seedMapLookup(walker, stf)
		if target_index != -1 {
			walker = target_index
			break
		}
	}
	// Now we're a fertilizer
	for _, ftw := range s.fertilizer_to_water {
		target_index := seedMapLookup(walker, ftw)
		if target_index != -1 {
			walker = target_index
			break
		}
	}

	// Now we're a water
	for _, wtl := range s.water_to_light {
		target_index := seedMapLookup(walker, wtl)
		if target_index != -1 {
			walker = target_index
			break
		}
	}

	// Now we're a light
	for _, ltt := range s.light_to_temperature {
		target_index := seedMapLookup(walker, ltt)
		if target_index != -1 {
			walker = target_index
			break
		}
	}

	// Now we're a temperature
	for _, tth := range s.temperature_to_humidity {
		target_index := seedMapLookup(walker, tth)
		if target_index != -1 {
			walker = target_index
			break
		}
	}

	// Now we're a humidity
	for _, htl := range s.humidity_to_location {
		target_index := seedMapLookup(walker, htl)
		if target_index != -1 {
			walker = target_index
			break
		}
	}
	return walker
}

func (s SeedInput) GetLocationsPartB() {
	// So what we need to do is mutate each map to the only reachable values, then repeat
}

func GetReachableLocations(currentPos int, startRange int, targetMap [][]int) [][]int {
	reachable_positions := [][]int{}
	to_travel := startRange

	// Because targetMap is sorted, we can rely on that feature
	for _, target := range targetMap {
		if to_travel == 0 {
			return reachable_positions
		}
		destinationStart := target[0]
		// destinationEnd := target[0] + target[2]
		sourceStart := target[1]
		// sourceEnd := target[1] + target[2]
		targetRange := target[2]

		if currentPos < sourceStart {
			continue
		}

		distance_from_source_start := currentPos - sourceStart
		if sourceStart > currentPos {
			edge_start := currentPos + distance_from_source_start
			edge_size := sourceStart - currentPos
			reachable_positions = append(reachable_positions, []int{edge_start, edge_size})
			to_travel -= edge_size
			if to_travel == 0 {
				return reachable_positions
			}
		}

		if distance_from_source_start > target[2] {
			continue
		}

		edge_start := destinationStart + distance_from_source_start
		edge_size := targetRange - distance_from_source_start
		currentPos += edge_size
		if edge_size > to_travel {
			edge_size = to_travel
			to_travel = 0
		} else {
			to_travel -= edge_size
		}

		reachable_positions = append(reachable_positions, []int{edge_start, edge_size})
	}
	if to_travel != 0 {
		reachable_positions = append(reachable_positions, []int{currentPos, to_travel})
	}
	return reachable_positions
}

func GetAllReachableLocations(inputs [][]int, targetMap [][]int) [][]int {
	reachable_locations := [][]int{}
	for _, input := range inputs {
		new_map := GetReachableLocations(input[0], input[1], targetMap)
		fmt.Println("Transformed: ", input, new_map)
		reachable_locations = append(reachable_locations, new_map...)
	}
	return reachable_locations
}

func sortArrayofArraysofInt(array [][]int) {
	sort.SliceStable(array, func(i, j int) bool {
		return array[i][1] < array[j][1]
	})

}

func parse_seeds(data []string) SeedInput {
	input := SeedInput{}
	state := 0
	skip_header := false
	for _, row := range data {
		if skip_header == true {
			skip_header = false
			continue
		}
		if row == "" {
			state += 1
			skip_header = true
			continue
		}
		if state == 0 {
			seeds := strings.Split(strings.Split(row, ":")[1], " ")
			input.seeds = toListOfInts(seeds)
			continue
		}
		row_split := strings.Split(row, " ")
		row_parsed := toListOfInts(row_split)
		if state == 1 {
			input.seed_to_soil = append(input.seed_to_soil, row_parsed)
		}
		if state == 2 {
			input.soil_to_fertilizer = append(input.soil_to_fertilizer, row_parsed)
		}
		if state == 3 {
			input.fertilizer_to_water = append(input.fertilizer_to_water, row_parsed)
		}
		if state == 4 {
			input.water_to_light = append(input.water_to_light, row_parsed)
		}
		if state == 5 {
			input.light_to_temperature = append(input.light_to_temperature, row_parsed)
		}
		if state == 6 {
			input.temperature_to_humidity = append(input.temperature_to_humidity, row_parsed)
		}
		if state == 7 {
			input.humidity_to_location = append(input.humidity_to_location, row_parsed)
		}
	}

	sortArrayofArraysofInt(input.seed_to_soil)
	sortArrayofArraysofInt(input.soil_to_fertilizer)
	sortArrayofArraysofInt(input.fertilizer_to_water)
	sortArrayofArraysofInt(input.water_to_light)
	sortArrayofArraysofInt(input.light_to_temperature)
	sortArrayofArraysofInt(input.temperature_to_humidity)
	sortArrayofArraysofInt(input.humidity_to_location)
	return input
}

func q5part1(seedInput SeedInput) int {
	lowest := 9999999999
	for i := 0; i < len(seedInput.seeds); i++ {
		location := seedInput.GetLocation(i)
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func q5part2(seedInput SeedInput) int {
	active_range := [][]int{}
	for i := 0; i < len(seedInput.seeds)/2; i++ {
		active_range = append(active_range, []int{seedInput.seeds[2*i], seedInput.seeds[2*i+1]})
	}

	active_range = GetAllReachableLocations(active_range, seedInput.seed_to_soil)
	active_range = GetAllReachableLocations(active_range, seedInput.soil_to_fertilizer)
	active_range = GetAllReachableLocations(active_range, seedInput.fertilizer_to_water)
	active_range = GetAllReachableLocations(active_range, seedInput.water_to_light)
	active_range = GetAllReachableLocations(active_range, seedInput.light_to_temperature)
	active_range = GetAllReachableLocations(active_range, seedInput.temperature_to_humidity)
	active_range = GetAllReachableLocations(active_range, seedInput.humidity_to_location)

	min_location := 999999999
	for _, i := range active_range {
		if i[0] < min_location {
			min_location = i[0]
		}
	}
	return min_location
}
