package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var INPUT string

func makeRow(line string) string {
	row := strings.TrimLeft(strings.ToLower(strings.Split(line, ":")[1]), " ")
	return row
}

func main() {
	data := strings.Split(strings.TrimRight(INPUT, "\n"), "\n")

	var sum int

	for _, row := range data {
		id, _ := strconv.Atoi(strings.TrimLeft(strings.Split(row, ":")[0], "Game "))
		plays := strings.TrimLeft(strings.Split(row, ":")[1], "Game ")
		if checkPlays(plays) {
			sum += id
		}
	}
	fmt.Println(sum)
}

func checkPlays(s string) bool {
	plays := strings.Split(s, "; ")

	for _, p := range plays {
		current := strings.Split(p, ", ")
		if !checkEachLine(current) {
			return false
		}
	}
	return true
}

func checkEachLine(s []string) bool {
	for _, row := range s {

		color_of_cube := strings.Split(row, " ")[1]
		fubar := strings.Split(row, " ")[0]
		num_of_cube, _ := strconv.Atoi(fubar)
		if color_of_cube == "red" && num_of_cube <= 12 || color_of_cube == "green" && num_of_cube <= 13 || color_of_cube == "blue" && num_of_cube <= 14 {
			continue
		} else {
			return false
		}
	}
	return true
}
