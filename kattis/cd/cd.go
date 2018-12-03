package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	max = 1000000
)

var scanner = bufio.NewScanner(os.Stdin)
var jack, jill [max]int

var line string
var numJack, numJill int

var cd int64

var count int
var x, y int
var space int
var i int

func main() {
	for {
		scanner.Scan()
		line = scanner.Text()

		for space = 0; space < len(line); space++ {
			if line[space] == 32 {
				break
			}
		}

		numJack, _ = strconv.Atoi(line[:space])
		numJill, _ = strconv.Atoi(line[space+1:])

		// fmt.Printf("numjack=%d numjill=%d\n", numJack, numJill)

		if numJack == 0 && numJill == 0 {
			break
		}

		for i = 0; i < numJack; i++ {
			scanner.Scan()
			cd, _ = strconv.ParseInt(scanner.Text(), 10, 32)
			jack[i] = int(cd)
		}

		for i = 0; i < numJill; i++ {
			scanner.Scan()
			cd, _ = strconv.ParseInt(scanner.Text(), 10, 32)
			jill[i] = int(cd)
		}

		// fmt.Println(jack[:numJack], jill[:numJill])

		count = 0
		x = 0
		y = 0

		for {
			if x >= numJack || y >= numJill {
				break
			} else if jack[x] == jill[y] {
				count++
				x++
				y++
			} else if jack[x] > jill[y] {
				y++
			} else {
				x++
			}
		}

		os.Stdout.Write([]byte(strconv.Itoa(count)))
		os.Stdout.Write([]byte{10})
	}
}
