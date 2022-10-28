package main

import (
	"fmt"
)

const OvenTime int = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}

func main() {
	actualTime := 10
	remainingTime := RemainingOvenTime(actualTime)
	fmt.Println("Remaining time", remainingTime)

}
