package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}

		}
	}
	return nil
}

func main() {
	var x []int = []int{1, 2, 3}

	tar := 4
	var result []int
	result = twoSum(x, tar)
	fmt.Println(result)

}
