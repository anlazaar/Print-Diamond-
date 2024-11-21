package main

import (
	"Diamond/utils"
	"os"
	"strconv"
)

func main() {
	if !utils.CheckAruments(os.Args) {
		utils.Usage()
		return
	}

	height, _ := strconv.Atoi(os.Args[1])
	width, _ := strconv.Atoi(os.Args[2])

	utils.PrintDiamond(height, width)
}
