package main

import (
	"fmt"
	"math"
	"math/rand"
)

type diceProfile [6]int
type results [3]int
type dicesRoll struct {
	diceProfile   // store if the face is failure (0), success (1) or epic fail (2)
	results       // store the number of success, failure, and epic
	total         int
	reroll        int
	resultPerCent [3]int
}

func main() {
	var myResults dicesRoll
	myResults.total = 5
	myResults.reroll = 0
	myResults.diceProfile[0] = 1
	myResults.diceProfile[1] = 1
	myResults.diceProfile[2] = 1
	myResults.diceProfile[3] = 1
	myResults.diceProfile[4] = 0
	myResults.diceProfile[5] = 2

	//throw total number of time
	for index := 0; index < myResults.total; index++ {
		randNum := rand.Intn(6)
		checkSuccess := myResults.diceProfile[randNum]
		resultWithRR := reroll(checkSuccess, myResults.reroll, myResults.diceProfile)
		myResults.results[resultWithRR]++
	}

	myResults.resultPerCent[0] = percent(myResults.results[0], myResults.total)
	myResults.resultPerCent[1] = percent(myResults.results[1], myResults.total)
	myResults.resultPerCent[2] = percent(myResults.results[2], myResults.total)

	fmt.Println(myResults.resultPerCent)

}
func percent(occurrence int, total int) int {
	floatResult := math.Round(100 * float64(occurrence) / float64(total))
	return int(floatResult)
}
func reroll(result int, nReroll int, table [6]int) int {
	var isEpic int
	if result == 1 {
		return 1
	}
	if result == 2 {
		isEpic = 1
	}

	// Rerolling for loop
	for index := 0; index < nReroll; index++ {
		fmt.Println("hello")
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
	return 0
}
