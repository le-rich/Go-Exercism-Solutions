package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, bird := range birdsPerDay {
		total += bird
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	subSlice := birdsPerDay[((week - 1) * 7):(((week - 1) * 7) + 7)]
	total := 0
	for _, bird := range subSlice {
		total += bird
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx, _ := range birdsPerDay {
		if idx%2 == 0 {
			birdsPerDay[idx]++
		}
	}
	return birdsPerDay
}
