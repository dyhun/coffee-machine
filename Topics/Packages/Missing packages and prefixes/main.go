package main

import (
	"fmt"
	. "fmt"
	"math"
	. "math"
)

func main() {
	// just add the missing prefixes to the functions below.
	var angle float64
	fmt.Scanf("%f", &angle)

	angle = angle * (math.Pi / 180)

	Println(Sin(angle) - math.Cos(angle))
}
