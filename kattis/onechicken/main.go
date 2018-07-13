package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const space = 32

var scanner *bufio.Scanner
var people, chicken, leftover int64

var bytes []byte
var lestring string
var spaceIndex int

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	bytes = scanner.Bytes()
	lestring = string(bytes)

	for i, b := range bytes {
		if b == space {
			spaceIndex = i
			break
		}
	}

	people, _ = strconv.ParseInt(string(bytes[:spaceIndex]), 10, 16)
	chicken, _ = strconv.ParseInt(string(bytes[spaceIndex+1:]), 10, 16)

	leftover = chicken - people

	if leftover == 1 {
		fmt.Println("Dr. Chaz will have 1 piece of chicken left over!")
	} else if leftover == -1 {
		fmt.Println("Dr. Chaz needs 1 more piece of chicken!")
	} else if leftover > 1 {
		fmt.Printf("Dr. Chaz will have %d pieces of chicken left over!\n", leftover)
	} else {
		fmt.Printf("Dr. Chaz needs %d more pieces of chicken!\n", -leftover)
	}
}
