package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	length := len(slice)
	if index >= length || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	length := len(slice)
	if index >= length || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newSlice := []int{}
	for _, v := range values {
		newSlice = append(newSlice, v)
	}
	return append(newSlice, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	length := len(slice)
	if index >= length || index < 0 {
		return slice
	}
	if index == length-1 {
		return append([]int{}, slice[:index]...)
	}
	newSlice := make([]int, len(slice[:index]))
	copy(newSlice, slice[:index])
	fmt.Println(newSlice)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}
func changeVal(s []int) {
	s[0] = 10
	s1 := s[:2]
	s1[0] = 90
}

func main() {
	s := []int{3, 4, 5, 6}
	fmt.Println(s[6])
}
