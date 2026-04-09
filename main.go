package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	argCount := len(args)

	//Strict usage check
	if argCount < 1 || argCount > 3 {
		printUsage()
		return
	}

	var outputFileName string
	var inputStr string
	var bannerName = "standard"

	//Check for optional project: --output flag
	if strings.HasPrefix(args[0], "--output=") {
		outputFileName = strings.TrimPrefix(args[0], "--output")

		//Invalid format if flag is empty or no string follows it
		if outputFileName == "" || argCount < 2 {
			printUsage()
			return
		}

		inputStr = args[1]
		if argCount == 3 {
			bannerName = args[2]
		}
	} else {
		//Stand mode & single argument mode
		inputStr = args[0]
		if argCount == 2 {
			bannerName = args[1]
		} else if argCount == 3 {
			//3 args provided but first isn't --output
			printUsage()
			return
		}
	}

	//Execution: Get the ASCII string
	result, err := GenerateASCII(inputStr, bannerName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	//FS API: Output Dispatcher
	if outputFileName != "" {
		//Write to file with 0644 permissions (standard for text files)
		err := os.WriteFile(outputFileName, []byte(result), 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	} else {
		//Print to terminal
		fmt.Print(result)
	}

}

func printUsage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
}
