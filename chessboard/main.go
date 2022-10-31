package main

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int = 0
	v, exist := cb[file]
	if !exist {
		return counter
	}
	for _, b := range v {
		if b {
			counter++
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var counter int = 0
	for _, file := range cb {
		for i, v := range file {
			if (i+1) == rank && v {
				counter++
			}
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var counter int = 0
	for _, v := range cb {
		counter += len(v)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int = 0
	for _, v := range cb {
		for _, b := range v{
			if b{
				counter += 1
			}
		}
	}
	return counter
}
