package main

import "fmt"

var original = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 10, 23, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 35,
	9, 39, 2, 39, 13, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 9, 59, 63, 2, 6, 63, 67, 1, 13, 67, 71, 1, 9, 71,
	75, 2, 13, 75, 79, 1, 79, 10, 83, 2, 83, 9, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 99, 5, 103, 1, 103, 2, 107, 1, 107, 10, 0, 99, 2, 0, 14, 0}

func main() {

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			original[1] = noun
			original[2] = verb
			var cop = make([]int, len(original))
			copy(cop, original)
			var result = calculate(cop)
			if result == 19690720 {
				fmt.Printf("\n Found it, noun: %d and verb %d", noun, verb)
				fmt.Printf("And the answer: %d", 100*noun+verb)
			}
		}
	}
}

func calculate(input []int) int {
	for i := 0; i < len(input); i += 4 {
		if input[i] == 99 {
			break
		}
		switch input[i] {
		case 1:
			if input[i+3] < len(input) && input[i+2] < len(input) && input[i+1] < len(input) {
				input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
			}
		case 2:
			if input[i+3] < len(input) && input[i+2] < len(input) && input[i+1] < len(input) {
				input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
			}
		case 99:
			break
		default:
			break
		}
	}
	return input[0]
}
