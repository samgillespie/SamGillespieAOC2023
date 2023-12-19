package answers

import (
	"strconv"
	"strings"
)

func Day19() []interface{} {
	data := ReadInputAsStr(19)
	workflows, parts := parse_workflows(data)
	return []interface{}{q19part1(workflows, parts), q19part2(workflows, parts)}
}

type WorkflowRule struct {
	list_element int
	condition    byte
	value        int
	target       string
}

func parse_workflows(data []string) (map[string][]WorkflowRule, [][]int) {
	parsing_rules := true
	workflows := map[string][]WorkflowRule{}
	parts := [][]int{}
	for _, row := range data {
		if parsing_rules == true {
			if row == "" {
				parsing_rules = false
				continue
			}
			row_split := strings.Split(row, "{")
			name := row_split[0]
			conditions := strings.Split(row_split[1][0:len(row_split[1])-1], ",")
			rules := []WorkflowRule{}
			for _, condition := range conditions {
				var rule WorkflowRule
				if strings.Contains(condition, ":") == false {
					rule = WorkflowRule{
						condition: 't', // t = true
						target:    condition,
					}
				} else {
					condition_split := strings.Split(condition, ":")
					value, err := strconv.Atoi(condition_split[0][2:])
					if err != nil {
						panic(err)
					}
					var list_element int
					switch {
					case condition[0] == 'x':
						list_element = 0
					case condition[0] == 'm':
						list_element = 1
					case condition[0] == 'a':
						list_element = 2
					case condition[0] == 's':
						list_element = 3
					default:
						panic("unexpected value" + condition)

					}
					rule = WorkflowRule{
						list_element: list_element,
						condition:    condition[1],
						value:        value,
						target:       condition_split[1],
					}
				}
				rules = append(rules, rule)
			}
			workflows[name] = rules
		} else {
			row = row[1 : len(row)-1]
			row_split := strings.Split(row, ",")
			x, _ := strconv.Atoi(row_split[0][2:])
			m, _ := strconv.Atoi(row_split[1][2:])
			a, _ := strconv.Atoi(row_split[2][2:])
			s, _ := strconv.Atoi(row_split[3][2:])
			parts = append(parts, []int{x, m, a, s})
		}
	}
	return workflows, parts
}

func process_workflow(workflows []WorkflowRule, part []int) string {
	for _, workflow := range workflows {
		if workflow.condition == 't' {
			return workflow.target
		}
		if workflow.condition == '>' {
			if part[workflow.list_element] > workflow.value {
				return workflow.target
			}
		}
		if workflow.condition == '<' {
			if part[workflow.list_element] < workflow.value {
				return workflow.target
			}
		}
	}
	panic("workflow failed")
}

func part_sum(part []int) int {
	return part[0] + part[1] + part[2] + part[3]
}

func q19part1(workflows map[string][]WorkflowRule, parts [][]int) int {
	total := 0
	for _, part := range parts {
		rule := "in"
		for {
			workflow, exists := workflows[rule]
			if !exists {
				panic("Could not find workflow " + rule)
			}
			rule = process_workflow(workflow, part)
			if rule == "A" {
				total += part_sum(part)
				break
			}
			if rule == "R" {
				break
			}
		}
	}
	return total
}

type WorkflowSubset struct {
	x_min     int
	x_max     int
	m_min     int
	m_max     int
	a_min     int
	a_max     int
	s_min     int
	s_max     int
	rule_name string
}

func newWorkflowSubset() WorkflowSubset {
	return WorkflowSubset{
		x_min: 1, m_min: 1, a_min: 1, s_min: 1,
		x_max: 4000, m_max: 4000, a_max: 4000, s_max: 4000,
		rule_name: "in"}
}

func (w WorkflowSubset) copy() WorkflowSubset {
	return WorkflowSubset{
		x_min: w.x_min,
		x_max: w.x_max,
		m_min: w.m_min,
		m_max: w.m_max,
		a_min: w.a_min,
		a_max: w.a_max,
		s_min: w.s_min,
		s_max: w.s_max,
	}
}

func (w WorkflowSubset) total_range() int {
	// Plus one because limits are inclusive
	return (w.x_max + 1 - w.x_min) * (w.m_max + 1 - w.m_min) * (w.a_max + 1 - w.a_min) * (w.s_max + 1 - w.s_min)
}

func process_rule_with_subset(subset WorkflowSubset, rules []WorkflowRule) ([]WorkflowSubset, int, int) {
	new_rules := []WorkflowSubset{}
	accepted := 0
	rejected := 0
	for _, rule := range rules {
		fail_condition := subset.copy()

		// if rule = value > 2001
		// Threshold 2002- 4000 and 1 - 2001 inclusive
		if rule.condition == '>' {
			switch {
			case rule.list_element == 0:
				subset.x_min = max(subset.x_min, rule.value+1)
				fail_condition.x_max = max(fail_condition.x_min, rule.value)
			case rule.list_element == 1:
				subset.m_min = max(subset.m_min, rule.value+1)
				fail_condition.m_max = max(fail_condition.m_min, rule.value)
			case rule.list_element == 2:
				subset.a_min = max(subset.a_min, rule.value+1)
				fail_condition.a_max = max(fail_condition.a_min, rule.value)
			case rule.list_element == 3:
				subset.s_min = max(subset.s_min, rule.value+1)
				fail_condition.s_max = max(fail_condition.s_min, rule.value)
			default:
				panic("unexpected")
			}
			subset.rule_name = rule.target
			if subset.rule_name == "A" {
				accepted += subset.total_range()
			} else if subset.rule_name == "R" {
				rejected += subset.total_range()
			} else {
				new_rules = append(new_rules, subset)
			}
			subset = fail_condition

		} else if rule.condition == '<' {
			// if rule = value < 2001
			// Threshold 1- 2000 and 2001 - 4000 inclusive
			switch {
			case rule.list_element == 0:
				subset.x_max = min(subset.x_max, rule.value-1)
				fail_condition.x_min = min(fail_condition.x_max, rule.value)
			case rule.list_element == 1:
				subset.m_max = min(subset.m_max, rule.value-1)
				fail_condition.m_min = min(fail_condition.m_max, rule.value)
			case rule.list_element == 2:
				subset.a_max = min(subset.a_max, rule.value-1)
				fail_condition.a_min = min(fail_condition.a_max, rule.value)
			case rule.list_element == 3:
				subset.s_max = min(subset.s_max, rule.value-1)
				fail_condition.s_min = min(fail_condition.s_max, rule.value)
			default:
				panic("unexpected")
			}
			subset.rule_name = rule.target
			if subset.rule_name == "A" {
				accepted += subset.total_range()
			} else if subset.rule_name == "R" {
				rejected += subset.total_range()
			} else {
				new_rules = append(new_rules, subset)
			}
			subset = fail_condition
		} else if rule.condition == 't' {
			if rule.target == "A" {
				accepted += subset.total_range()
			} else if rule.target == "R" {
				rejected += subset.total_range()
			} else {
				subset.rule_name = rule.target
				new_rules = append(new_rules, subset)
			}
		}
	}
	return new_rules, accepted, rejected
}

func q19part2(workflows map[string][]WorkflowRule, parts [][]int) int {
	active_subsets := []WorkflowSubset{}
	active_subsets = append(active_subsets, newWorkflowSubset())
	accepted := 0
	for len(active_subsets) > 0 {
		next_iteration := []WorkflowSubset{}
		for _, subset := range active_subsets {
			workflow := workflows[subset.rule_name]
			subsets, accept, _ := process_rule_with_subset(subset, workflow)
			accepted += accept
			next_iteration = append(next_iteration, subsets...)

		}
		active_subsets = next_iteration
	}
	return accepted
}
