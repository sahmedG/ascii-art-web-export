package webart

import (
	"bufio"
	"fmt"
	"os"
)

func PrintFileLine(lineNumber int, filePath string) string {
	//* this program scans the banner to get the art
	file, err := os.Open(filePath)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file: ", err, ", please enter a valid banner file name")
		os.Exit(1)
	}
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Read line by line
	line := ""
	lineCount := 0 // A counter used to stop at specific line
	for scanner.Scan() {
		lineCount++
		// save the line and stop the loop
		if lineCount == lineNumber {
			line = scanner.Text()
			break
		}
	}
	// fmt.Print(line)
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return ""
	}
	return line
}
