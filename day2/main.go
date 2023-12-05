package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Game map[string]int
type Games []map[string]int
type Set []string

//go:embed input.txt
var INPUT string

func main() {
	partOne()
	partTwo()
}

func createSets() Set {
	var sets Set
	data := strings.Split(strings.TrimRight(INPUT, "\n"), "\n")
	for _, n := range data {
		test := strings.Split(n, ": ")
		sets = append(sets, test[1])
	}
	return sets
}

func getListOfAllGame(data Set) Games {
	var games Games

	for _, row := range data {
		game := make(map[string]int)
		for _, sets := range strings.Split(row, "; ") {
			for _, individual := range strings.Split(sets, ", ") {
				color := strings.Split(individual, " ")[1]
				num, _ := strconv.Atoi(strings.Split(individual, " ")[0])

				if _, ok := game[color]; ok {
					if game[color] < num {
						game[color] = num
					}
				} else {
					game[color] = num
				}
			}
		}
		games = append(games, game)
	}

	return games
}

func partTwo() {

	data := createSets()
	games := getListOfAllGame(data)
	calcGames(games)
}

func calcGames(g Games) {

	var sum int

	for _, n := range g {
		total := n["red"] * n["blue"] * n["green"]
		sum += total
	}

	fmt.Println(sum)

}

func partOne() {
	var sum int
	for _, row := range strings.Split(strings.TrimRight(INPUT, "\n"), "\n") {
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
		num_of_cube, _ := strconv.Atoi(strings.Split(row, " ")[0])

		if color_of_cube == "red" && num_of_cube > 12 || color_of_cube == "green" && num_of_cube > 13 || color_of_cube == "blue" && num_of_cube > 14 {
			return false
		}
	}
	return true
}
