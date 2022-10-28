package main

import "fmt"

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate) * successRate) / 100 / 60)
}

func CalculateCost(carsCount int) uint {
	group_of_ten := carsCount / 10
	individual := carsCount - (group_of_ten * 10)
	total := (group_of_ten * 95000) + (individual * 10000)
	return uint(total)
}

func main() {
	// result1 := CalculateWorkingCarsPerHour(1547, 90)
	// result2 := CalculateWorkingCarsPerMinute(1105, 90)
	result3 := CalculateCost(21)
	fmt.Println(result3)
}
