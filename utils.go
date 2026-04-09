package main

// GetLineIndex calculates the starting line of a character row in the banner file.
// Linear Offset Formula: (CharIndex * BlockHeight) + RowOffset
func GetLineIndex(char rune, row int) int {
	//Printable ASCII starts at Space (32) and ends at Tilde (126)
	if char < 32 || char > 126 {
		char = '?' // Fallback for unprintable characters
	}

	//Each character block is 9 lines tall (8 art lines + 1 separator)
	return int(char-32)*9 + row

}

// IsOnlyNewLines determine if the input contains text or just \n
func IsOnlyNewLines(lines []string) bool {
	for _, l := range lines {
		if l != "" {
			return false
		}
	}
	return len(lines) > 1
}
