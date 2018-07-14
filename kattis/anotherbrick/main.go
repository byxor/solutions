package main

import (
	"fmt"
)

func main() {
	var height, width, numBricks int
	fmt.Scanf("%d %d %d\n", &height, &width, &numBricks)

	bricks := make([]int, numBricks)
	ReadN(bricks, 0, numBricks)

	if width == 0 || height == 0 {
		fmt.Println("YES")
	}

	usedWidth := 0
	usedHeight := 0

	for _, brick := range bricks {
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

	if usedWidth == 0 && usedHeight == height {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}
