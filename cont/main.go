package main

import "fmt"

func flipString(str string) string {
	flipped := []rune(str)
	fmt.Println(flipped)
	for i, char := range str {
		switch char {
		case 'a':
			flipped[i] = 'z'
		default:
			flipped[i] = char + 1
		}
	}
	return string(flipped)
}
func main() {
	fmt.Println(flipString("abiolaABCD1234"))
}
