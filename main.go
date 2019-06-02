package main

import (
	"fmt"
	"math"
	"math/rand"
)

type faceCount [6]int
type results [3]int
type failOrSucess [6]int
type dicesRoll struct {
	faceCount
	failOrSucess
	results
	total         int
	reroll        int
	resultPerCent [3]float64
}

func main() {
	var myResults dicesRoll
	myResults.total = 1000
	myResults.failOrSucess[0] = 1
	myResults.failOrSucess[1] = 1
	myResults.failOrSucess[5] = 2

	for index := 0; index < myResults.total; index++ {
		randNum := rand.Intn(6)
		myResults.faceCount[randNum]++
		if myResults.failOrSucess[randNum] == 0 {
			myResults.results[0]++
		}
		if myResults.failOrSucess[randNum] == 1 {
			myResults.results[1]++
		}
		if myResults.failOrSucess[randNum] == 2 {
			myResults.results[2]++
		}
	}

	myResults.resultPerCent[0] = math.Round(100 * float64(myResults.results[0]) / float64(myResults.total))
	myResults.resultPerCent[1] = math.Round(100 * float64(myResults.results[1]) / float64(myResults.total))
	myResults.resultPerCent[2] = math.Round(100 * float64(myResults.results[2]) / float64(myResults.total))

	fmt.Println(myResults.resultPerCent)

}
