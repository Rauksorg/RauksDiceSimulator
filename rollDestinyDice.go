package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
refactor
create setup function
resulttoPercent -> function
throw dice -> function
*/

func destiny(numberOfDice int, diceProfile [6]int, choose string) int {
	arrayResult := make([]int, numberOfDice)
	for index := 0; index < numberOfDice; index++ {
		randNum := rand.Intn(6)
		arrayResult[index] = randNum
	}
	sort.Ints(arrayResult)
	best := arrayResult[len(arrayResult)-1]
	worst := arrayResult[0]
	if choose == "worst" {
		return worst
	}
	return best
}

func rollDestinyDice() {
	type diceProfile [6]int
	type destinyDice struct {
		diceProfile  diceProfile // store if the face is failure (0), success (1) or epic fail (2)
		results      [4]int      // store the number of success, failure, and epic
		total        int
		numberOfDice int
		choose       string
	}

	myResults2 := destinyDice{
		total:        1000000,
		numberOfDice: 2,
		diceProfile:  diceProfile{0, 0, 1, 1, 2, 2},
		choose:       "worst",
	}

	//throw destiny dice total number of time
	for index := 0; index < myResults2.total; index++ {
		destinyResult := destiny(myResults2.numberOfDice, myResults2.diceProfile, myResults2.choose)
		checkSuccess := myResults2.diceProfile[destinyResult]
		myResults2.results[checkSuccess]++
	}
	fmt.Println("No :", toPercent(myResults2.results[0], myResults2.total), "| Neutral :", toPercent(myResults2.results[1], myResults2.total), "| Yes :", toPercent(myResults2.results[2], myResults2.total), "| Epic :", toPercent(myResults2.results[3], myResults2.total))

}
