package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Data []string

//go:embed input.txt
var input string

var dict = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	partOne()

}

func partOne() {

	rows := parseData(input)
	// rows.Print()

	for _, n := range rows {

		getResult(n)

	}
}

func checkRow(row string) int {

	if unicode.IsDigit(rune(row[0])) {
		n, _ := strconv.Atoi(string(row[0]))
		return n
	}

	for key, val := range dict {
		if strings.HasPrefix(row, val) {
			return key
		}

	}

	return -1
}

func getResult(row string) string {

	var result string

	first := 0
	last := 0
	for _, row := range strings.Split(row, "") {
		row_length := len(row)

		for i := 0; i < row_length; i++ {
			digit := checkRow(row[i:])
			if digit != -1 {
				if first == 0 {
					first = digit
				}
				last = digit

			}
		}

	}

	result = fmt.Sprintf("%d%d", first, last)

	fmt.Println(result)

	return result
}

func parseData(s string) Data {
	a := strings.TrimRight(s, "\n")
	b := strings.Split(a, "\n")
	return b

}

func (s *Data) Print() {
	for _, row := range *s {
		fmt.Println(row)
	}

}
