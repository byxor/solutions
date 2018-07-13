package main

import (
	"fmt"
	"math"
)

func main() {
	var matches int
	var width, height float64
	fmt.Scanf("%d %f %f\n", &matches, &width, &height)
	boxLength := hyp(width, height)

	for i := 0; i < matches; i++ {
		var matchLength int
		fmt.Scanf("%d\n", &matchLength)
		if float64(matchLength) <= boxLength {
			fmt.Println("DA")
		} else {
			fmt.Println("NE")
		}
	}
}

func hyp(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
