package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("./quick_sort_input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	input := make([]int, 0)
	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		input = append(input, num)
	}
	fmt.Println(input)
}


