package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var INPUT string

func makeData() [][]string {
	step_one := strings.Split(strings.TrimRight(INPUT, "\n"), "\n")
	data := returnOnlyNumbersFromRows(step_one)
	return data
}

func returnOnlyNumbersFromRows(rows []string) [][]string {
	var int_data [][]string

	for _, el := range rows {
		var line []string
		for _, x := range strings.Split(el, "") {
			_, err := strconv.Atoi(x)
			if err != nil {
				continue
			}
			line = append(line, x)
		}
		int_data = append(int_data, line)
	}

	return int_data
}

func partOne(data [][]string) {
	var sum int

	for _, n := range data {

		number_string := n[0] + n[len(n)-1]

		num, err := strconv.Atoi(number_string)
		if err != nil {
			fmt.Println("Fubar!!!!")
		}
		sum += num

	}
	fmt.Println(sum)
}

func main() {
	data := makeData()

	partOne(data)
}
