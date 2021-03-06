package main

import (
	"fmt"
)

var height, width, numBricks int
var usedWidth, usedHeight int
var brick int

func main() {
	fmt.Scan(&height)
	fmt.Scan(&width)
	fmt.Scan(&numBricks)

	for i := 0; i < numBricks; i++ {
		fmt.Scan(&brick)
		usedWidth += brick // lay a brick

		if usedWidth == width { // row is complete
			usedWidth = 0
			usedHeight++
		} else if usedWidth > width { // row has overflown
			fmt.Println("NO")
			return
		}

		if usedHeight == height { // height has been reached
			fmt.Println("YES")
			return
		}
	}
}
