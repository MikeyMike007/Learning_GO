package main

import (
		"fmt"
)

func add(x int, y int) int { // Could also be specified as add(x, y int) int
		return x + y
}



func main() {
		fmt.Println(add(42,17))
}
