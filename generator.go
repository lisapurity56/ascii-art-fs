package main

import (
	"os"
	"strings"
)

// GenerateASCII maps input characters to their 8-line banner reoresentations
func GenerateASCII(inputStr, bannerName string) (string, error) {
	//FS API Read the template file
	fileName := strings.TrimSuffix(bannerName, ".txt") + ".txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	//Data Manipulation Normalize line endings
	bannerLines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	//Handle literal \n and split the input into lines
	input := strings.ReplaceAll(inputStr, "\\n", "\n")
	inputLines := strings.Split(input, "\n")

	var builder strings.Builder

	// Egde Case: Check if the entire input is just newlines
	if IsOnlyNewLines(inputLines) {
		for i := 0; i < len(inputLines)-1; i++ {
			builder.WriteString("\n")
		}
		return builder.String(), nil
	}

	// Build the output (Horizontal assembly within 8 vertical lines)
	for _, word := range inputLines {
		if word == "" {
			builder.WriteString("\n")
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, char := range word {
				index := GetLineIndex(char, i)
				if index < len(bannerLines) {
					builder.WriteString(bannerLines[index])
				}
			}
			builder.WriteString("\n")
		}
	}
	return builder.String(), nil
}
