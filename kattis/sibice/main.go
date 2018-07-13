package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var matches int
	var width, height float64
	fmt.Scanf("%d %f %f\n", &matches, &width, &height)
	boxLength := width*width + height*height

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < matches; i++ {
		scanner.Scan()
		line := scanner.Text()
		matchLength, _ := strconv.Atoi(line)
		if float64(matchLength*matchLength) <= boxLength {
			fmt.Println("DA")
		} else {
			fmt.Println("NE")
		}
	}
}
