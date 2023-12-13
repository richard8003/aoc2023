package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type HighScores struct {
	TIMES     []int
	DISTANCES []int
}

type Race struct {
	TIME        int
	BUTTON_HOLD int
	DISTANCE    int
}

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne(parseData(input)))
}

func parseData(input string) HighScores {
	rows := strings.Split(strings.TrimRight(input, "\n"), "\n")
	times := stringSliceToIntSlice(strings.Fields(strings.Join(strings.Split(rows[0], ": "), " "))[1:])
	distnces := stringSliceToIntSlice(strings.Fields(strings.Join(strings.Split(rows[1], ": "), " "))[1:])

	return HighScores{
		TIMES:     times,
		DISTANCES: distnces,
	}

}
func stringSliceToIntSlice(s []string) []int {
	var ints []int
	for _, num := range s {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Something went wrong when converting this string to a number!\n")
		}
		ints = append(ints, n)
	}
	return ints
}

func checkEachRace(hs HighScores) int {
	var fubar []int

	for i := 0; i < len(hs.DISTANCES); i++ {
		distance := hs.DISTANCES[i]
		time := hs.TIMES[i]
		fubar = append(fubar, calcRace(time, distance))
	}

	// multiply all the answers
	sum := fubar[0]
	for i := 1; i < len(fubar); i++ {
		sum *= fubar[i]
	}

	return sum
}

func partOne(hs HighScores) string {
	sum := checkEachRace(hs)

	return fmt.Sprintf("Part 1: %d", sum)
}

func calcRace(raceTime int, oldRecord int) int {
	var sum int

	for i := 0; i <= raceTime; i++ {
		result, _ := calculateDistance(Race{
			TIME:        raceTime,
			BUTTON_HOLD: i,
		})
		if result > oldRecord {
			sum += 1
		}
	}
	return sum
}

func calculateDistance(r Race) (int, string) {
	r.DISTANCE = (r.TIME - r.BUTTON_HOLD) * r.BUTTON_HOLD
	return r.DISTANCE, fmt.Sprintf("Distance: %d", r.DISTANCE)
}
