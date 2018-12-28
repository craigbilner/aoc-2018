package main

import (
	"fmt"
)

func interpolate(generations int) int {
	return 5564 + ((generations - 97) * 40)
}

func main() {
	fmt.Printf("answer %v\n", interpolate(50000000000))
}
