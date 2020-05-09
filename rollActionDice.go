package main

import (
	"fmt"
	"math/rand"
)

/*
refactor
create setup function
resulttoPercent -> function
throw dice -> function
*/

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

func rollActionDice() {
	type diceProfile [6]int
	type actionDice struct {
		diceProfile diceProfile // store if the face is failure (0), success (1) or epic fail (2)
		results     [3]int      // store the number of success, failure, and epic
		total       int
		reroll      int
	}

	myResults := actionDice{
		total:       1000000,
		reroll:      0,
		diceProfile: diceProfile{1, 1, 0, 0, 0, 2},
	}

	//throw action dice total number of time
	for index := 0; index < myResults.total; index++ {
		randNum := rand.Intn(6)
		checkSuccess := myResults.diceProfile[randNum]
		resultWithRR := reroll(checkSuccess, myResults.reroll, myResults.diceProfile)
		myResults.results[resultWithRR]++
	}
	fmt.Println("Fail :", toPercent(myResults.results[0], myResults.total), "| Success :", toPercent(myResults.results[1], myResults.total), "| Epic fail :", toPercent(myResults.results[2], myResults.total))
}
