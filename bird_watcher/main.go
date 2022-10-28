package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for _, v := range birdsPerDay {
		count += v
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sliceStart int = 7 * (week - 1)
	var sliceStop int = sliceStart + 7
	return TotalBirdCount(birdsPerDay[sliceStart:sliceStop])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = v + 1
			continue
		}
	}
	return birdsPerDay
}

func main() {
	birds := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	total := TotalBirdCount(birds)
	fmt.Println(total)
	totalInWeek := BirdsInWeek(birds, 2)
	fmt.Println(totalInWeek)
	fmt.Println(FixBirdCountLog([]int{2, 5, 0, 7, 4, 1}))
}
