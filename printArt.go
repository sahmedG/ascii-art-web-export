package webart

import (
	"fmt"
	"strings"
)

//* this function prints the art while handling newlines

func PrintART(str string, fontname string) string {
	//* in the case newlines and only newlines were detected, the program prints a new line and exits
	str_word := ""
	if str == "\n" {
		fmt.Println()
		str_word += "\n"
		return ""
	}
	input_strs := strings.Split(str, "\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Println()
			str_word += "\n"
		} else {
			str_word += Print_Each_Rune_Line(word, fontname)
		}
	}
	return str_word
}
