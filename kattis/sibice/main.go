package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numMatches int
var width, height float64

var boxLength float64
var matchLength float64

var scanner *bufio.Scanner
var line string

func main() {
	fmt.Scanf("%d %f %f\n", &numMatches, &width, &height)
	boxLength = width*width + height*height
	scanner = bufio.NewScanner(os.Stdin)

	for i := 0; i < numMatches; i++ {
		scanner.Scan()
		line = scanner.Text()
		matchLength, _ = strconv.ParseFloat(line, 10)
		if float64(matchLength*matchLength) <= boxLength {
			fmt.Println("DA")
		} else {
			fmt.Println("NE")
		}
	}
}
