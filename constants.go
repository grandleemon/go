package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	fmt.Println(s)
	const a = 1000000000

	const b = 2e20 / a
	fmt.Println(b)

	fmt.Println(int64(b))

	fmt.Println(math.Sin(a))
}
