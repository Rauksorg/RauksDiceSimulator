package main

import (
	"fmt"
	"math"
	"math/rand"
)

type results [3]int
type failOrSuccess [6]int
type dicesRoll struct {
	failOrSuccess // store if the face is failure (0), success (1) or epic fail (2)
	results       // store the number of success, failure, and epic
	total         int
	reroll        int
	resultPerCent [3]float64
}

func main() {
	var myResults dicesRoll
	myResults.total = 1000
	myResults.reroll = 1
	myResults.failOrSuccess[0] = 1
	myResults.failOrSuccess[1] = 0
	myResults.failOrSuccess[5] = 2

	for index := 0; index < myResults.total; index++ {
		randNum := rand.Intn(6)
		checkSuccess := myResults.failOrSuccess[randNum]

		resultWithRR := reroll(checkSuccess, myResults.reroll, myResults.failOrSuccess)

		myResults.results[resultWithRR]++

	}

	myResults.resultPerCent[0] = math.Round(100 * float64(myResults.results[0]) / float64(myResults.total))
	myResults.resultPerCent[1] = math.Round(100 * float64(myResults.results[1]) / float64(myResults.total))
	myResults.resultPerCent[2] = math.Round(100 * float64(myResults.results[2]) / float64(myResults.total))

	fmt.Println(myResults.resultPerCent)

}
func reroll(result int, nReroll int, table [6]int) int {
	var isEpic int
	var newResult int
	if result == 2 {
		isEpic = 1
	}

	// If no reroll allowed
	if nReroll == 0 {
		if result == 1 {
			return 1
		}
		return 0
	}

	// Rerolling for loop
	for index := 0; index < nReroll; index++ {
		randNum := rand.Intn(6)
		if table[randNum] == 1 {
			return 1
		}
		if isEpic == 1 && table[randNum] != 1 {
			return 2
		}
		if table[randNum] == 2 {
			isEpic = 1
		}

	}
	return newResult
}
