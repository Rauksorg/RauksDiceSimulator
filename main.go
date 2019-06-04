package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

/*
refactor
create setup function
resulttoPercent -> function
throw dice -> function
*/
type diceProfile [6]int
type actionDice struct {
	diceProfile diceProfile // store if the face is failure (0), success (1) or epic fail (2)
	results     [3]int      // store the number of success, failure, and epic
	total       int
	reroll      int
}

type destinyDice struct {
	diceProfile  diceProfile // store if the face is failure (0), success (1) or epic fail (2)
	results      [5]int      // store the number of success, failure, and epic
	total        int
	numberOfDice int
}

func main() {

	myResults := actionDice{
		total:       1000000,
		reroll:      1,
		diceProfile: diceProfile{1, 0, 0, 0, 0, 2},
	}
	myResults2 := destinyDice{
		total:        10000000,
		numberOfDice: 4,
		diceProfile:  diceProfile{0, 0, 1, 1, 2, 3},
	}

	//throw action dice total number of time
	for index := 0; index < myResults.total; index++ {
		randNum := rand.Intn(6)
		checkSuccess := myResults.diceProfile[randNum]
		resultWithRR := reroll(checkSuccess, myResults.reroll, myResults.diceProfile)
		myResults.results[resultWithRR]++
	}
	fmt.Println("Fail :", toPercent(myResults.results[0], myResults.total), "| Success :", toPercent(myResults.results[1], myResults.total), "| Epic fail :", toPercent(myResults.results[2], myResults.total))

	//throw destiny dice total number of time
	for index := 0; index < myResults2.total; index++ {
		destinyResult := destiny(myResults2.numberOfDice, myResults2.diceProfile)
		myResults2.results[destinyResult]++
	}
	fmt.Println("0 :", toPercent(myResults2.results[0], myResults2.total), "| 1 :", toPercent(myResults2.results[1], myResults2.total), "| 2 :", toPercent(myResults2.results[2], myResults2.total), "| 3 :", toPercent(myResults2.results[3], myResults2.total), "| Epic :", toPercent(myResults2.results[4], myResults2.total))

}

func toPercent(occurrence int, total int) int {
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

func destiny(numberOfDice int, diceProfile [6]int) int {
	result := []int{}
	epic := 0
	for index := 0; index < numberOfDice; index++ {
		randNum := rand.Intn(6)
		checkSuccess := diceProfile[randNum]
		if checkSuccess == 3 {
			epic++
		}
		result = append(result, checkSuccess)
	}
	if epic >= 2 {
		return 4
	}
	sort.Ints(result)
	best := result[len(result)-1]
	return best
}
