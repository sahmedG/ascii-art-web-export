package webart


/*

this function maps every ascii number from 0 to 126 to the respective beggining line number in the banner files

*/


func MapART(char rune) int {
	return 9*(int(char)-32) + 2
}
