package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	args := os.Args
	var err error = errors.New("An error")
	k := 1
	var n, sum float64
	for err != nil {
		if k >= len(args) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(args[k], 64)
		k++
	}
	for i := 1; i < len(args); i++ {
		n, err = strconv.ParseFloat(args[i], 64)
		if err == nil {
			sum += n

		}
	}
	fmt.Println(sum)
}
