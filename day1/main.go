package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var INPUT string

func main() {
	data := parse_input_file()
	// partOne(data)
	partTwo(data)

}

func partTwo(data []string) {
	numberStrings := []string{
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
	for _, row := range data {

		first_digit := ""
		last_digit := ""

		for i := 0; i < len(row); i++ {
			if unicode.IsDigit(rune(row[i])) {
				if first_digit == "" {
					first_digit = string(row[i])
				}
				last_digit = string(row[i])
			}
			for _, nsr := range numberStrings {
				if strings.HasSuffix(row[:i], nsr) {
					first_digit = nsr
				}
			}
		}

		test := fmt.Sprintf("%s", first_digit+last_digit)
		fmt.Println(test)
	}

}

// func partOne(data []string) {
//
// // numberStrings := []string{
// // "one",
// // "two",
// // "three",
// // "four",
// // "five",
// // "six",
// // "seven",
// // "eight",
// // "nine",
// // }
// //
//
// var rows []int
//
// var sum int
//
// for _, row := range data {
//
// first := ""
// last := ""
//
// for i := 0; i < len(row); i++ {
// if unicode.IsDigit(rune(row[i])) {
// // fmt.Println(string(row[i]))
// if first == "" {
// first = string(row[i])
// }
// last = string(row[i])
// }
// }
//
// total, _ := strconv.Atoi(fmt.Sprintf("%s", first+last))
// rows = append(rows, total)
//
// sum += total
// }
// fmt.Println(rows)
// fmt.Println(sum)
// }
func parse_input_file() []string {
	result := strings.Split(strings.TrimRight(INPUT, "\n"), "\n")
	return result
}
