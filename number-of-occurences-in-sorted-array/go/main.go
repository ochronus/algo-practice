package main

import "fmt"

type result struct {
	result int
	steps  int
}

func naive(slice []int, lookingFor int) result {
	count := 0
	steps := 0
	for _, elem := range slice {
		steps = steps + 1
		if elem == lookingFor {
			count = count + 1
		}
	}
	return result{
		result: count,
		steps:  steps,
	}
}

func optim1(slice []int, lookingFor int) result {
	count := 0
	steps := 0
	for _, elem := range slice {
		steps = steps + 1
		if elem == lookingFor {
			count = count + 1
		} else {
			if elem > lookingFor {
				break
			}
		}
	}
	return result{
		result: count,
		steps:  steps,
	}
}

func binarySearch(slice []int, lookingFor int) result {

	start_index := 0
	end_index := len(slice) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if slice[median] < lookingFor {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	if start_index == len(slice) || slice[start_index] != lookingFor {
		return result{
			result: -1,
			steps:  1,
		}
	} else {
		return result{
			result: start_index,
			steps:  1,
		}
	}

}

func main() {
	haystack := []int{1, 2, 3, 4, 4, 4, 4, 5, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	needle := 4
	fmt.Printf("Searching for %d in %v\n", needle, haystack)
	fmt.Printf("Naive: %v\n", naive(haystack, needle))
	fmt.Printf("1st optimization: %v\n", optim1(haystack, needle))
}
