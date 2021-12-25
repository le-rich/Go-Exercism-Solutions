package reverse

func Reverse(input string) string {
	runeCopy := []rune(input)
	for i := 0; i < len(runeCopy) / 2; i++{
		temp := runeCopy[i]
		runeCopy[i] = runeCopy[len(runeCopy) - 1 - i]
		runeCopy[(len(runeCopy) - 1 - i)] = temp
	}

	return string(runeCopy)
}
