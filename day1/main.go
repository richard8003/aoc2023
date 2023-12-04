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
	// data := makeData()

	// partOne(data)

	data := strings.Split(strings.TrimRight(INPUT, "\n"), "\n")

	// fubar := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	// fmt.Println(fubar)

	fmt.Println(data)

	// for _, row := range data {
	// fmt.Println(row)
	// row = strings.ReplaceAll(row, "one", "1")
	// row = strings.ReplaceAll(row, "two", "2")
	// row = strings.ReplaceAll(row, "three", "3")
	// row = strings.ReplaceAll(row, "four", "4")
	// row = strings.ReplaceAll(row, "five", "5")
	// row = strings.ReplaceAll(row, "six", "6")
	// row = strings.ReplaceAll(row, "seven", "7")
	// row = strings.ReplaceAll(row, "eight", "8")
	// row = strings.ReplaceAll(row, "nine", "9")
	// row = strings.ReplaceAll(row, "zero", "0")
	// fmt.Println(row)
	// }

	// fmt.Println(data)

	// test := "jagHarEn"
	//
	// test = strings.ReplaceAll(test, "Har", "BAJS")
	// test = strings.ReplaceAll(test, "jag", "Du")
	//
	// fmt.Println(test)
}
