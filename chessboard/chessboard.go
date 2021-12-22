package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	if _, ok := cb[rank]; !ok {
		return 0
	} else {
		result := 0
		for _, val := range cb[rank] {
			if val {
				result += 1
			}
		}

		return result
	}
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	result := 0
	for _, element := range cb {
		if file < 0 || file-1 >= len(element) {
			return 0
		} else {
			if element[file-1] {
				result += 1
			}
		}
	}

	return result
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	result := 0
	for _, element := range cb {
		result += len(element)
	}

	return result
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	result := 0
	for _, element := range cb {
		for _, val := range element {
			if val {
				result++
			}
		}
	}
	return result
}
