package utils

import (
	"fmt"
	"strconv"
)

func Usage() {
	fmt.Println("[USAGE]: go run main.go [Diamond height] [Diamond width]")
}

// This Func will check the Arguments if there are valide type and count
func CheckAruments(args []string) bool {
	if len(args) != 3 {
		fmt.Println("Error: Invalide number of arguments.")
		return false
	}

	_, err1 := strconv.Atoi(args[1])
	_, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Height and width must be odd numbers to form a proper diamond.")
		return false
	}
	return true
}

func PrintDiamond(h, w int) {
	if h%2 == 0 || w%2 == 0 {
		fmt.Println("Error: Height and width must be odd numbers to form a proper diamond.")
		return
	}

	for i := 0; i < h; i++ {
		numStars := w - 2*abs(h/2-i)

		spaces := (w - numStars) / 2
		fmt.Print(repeatChar(' ', spaces))

		fmt.Print(repeatChar('*', numStars))

		fmt.Println()
	}
}

func abs(h int) int {
	if h < 0 {
		h = -h
	}
	return h
}

func repeatChar(char rune, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += string(char)
	}
	return result
}
