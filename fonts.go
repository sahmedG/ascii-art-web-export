package webart

import (
	"fmt"
	"os"
)

func MapFont(fontname string) string {
	var fonts = map[string]string{
		"standard":   "../standard.txt",
		"shadow":     "../shadow.txt",
		"thinkertoy": "../thinkertoy.txt",
	}
	_, err := os.Open(fonts[fontname])
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file: ", err, ", please enter a valid banner file name")
		return "Error 500"
	}
	return fonts[fontname]
}
