package main

import "math"

// go run main.go rollActionDice.go rollDestinyDice.go

func toPercent(occurrence int, total int) int {
	floatResult := math.Round(100 * float64(occurrence) / float64(total))
	return int(floatResult)
}
func main() {
	rollActionDice()
	rollDestinyDice()
}
