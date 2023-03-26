package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw_input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(raw_input), "\n\n")
	var max int

	for _, v := range input {
		calories := strings.Fields(v)
		var sum int

		for _, v := range calories {
			calory, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sum += calory
		}

		if sum > max {
			max = sum
		}

	}
	fmt.Println(max)
}
