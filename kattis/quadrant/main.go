package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner
var x, y int64

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	y, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	if x > 0 {
		if y > 0 {
			fmt.Println("1")
		} else {
			fmt.Println("4")
		}
	} else {
		if y > 0 {
			fmt.Println("2")
		} else {
			fmt.Println("3")
		}
	}

}
