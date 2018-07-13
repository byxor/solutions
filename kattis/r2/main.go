package main

import (
	"fmt"
)

var r1, mean int

func main() {
	fmt.Scanf("%d %d\n", &r1, &mean)
	fmt.Println(mean + mean - r1)
}
