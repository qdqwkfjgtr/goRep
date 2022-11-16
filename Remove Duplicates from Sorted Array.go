package main

import "fmt"

func removeDuplicates(nums []int) (int, []int) {
	p := nums[0]
	counter := 1
	var x []int
	if len(nums) == 0 {
		return 0, nums
	}
	x = append(x, p)
	for i := range nums {
		if p != nums[i] {
			p = nums[i]
			x = append(x, p)
			counter++

		}

	}
	return counter, x
}
func main() {
	arr := []int{1, 1, 2, 2, 4, 4, 5, 5, 5, 6, 6}
	fmt.Println(removeDuplicates(arr))
}
