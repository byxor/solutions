package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var line string
		var numJack, numJill int

		line, _ = reader.ReadString('\n')
		fmt.Sscanf(line, "%d %d", &numJack, &numJill)
		// fmt.Println(numJack, numJill)

		if numJack == 0 && numJill == 0 {
			break
		}

		jack := make([]int, numJack)
		readCds(reader, jack, numJack)

		jill := make([]int, numJill)
		readCds(reader, jill, numJill)

		// fmt.Println(jack, jill)

		var count int
		var x, y int

		var a, b int

		for {
			if x >= numJack || y >= numJill {
				break
			}

			a = jack[x]
			b = jill[y]

			// fmt.Printf("x=%d y=%d a=%d b=%d c=%d\n", x, y, a, b, count)
			if a == b {
				count++
				x++
				y++
			} else if a > b {
				y++
			} else {
				x++
			}
		}

		fmt.Println(count)
	}
}

func readCds(reader *bufio.Reader, cds []int, num int) {
	for i := 0; i < num; i++ {
		line, _ := reader.ReadString('\n')
		cd, _ := strconv.ParseInt(strings.TrimSuffix(line, "\n"), 10, 64)
		cds[i] = int(cd)
	}
}
