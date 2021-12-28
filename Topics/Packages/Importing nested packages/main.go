package main

import (
    "fmt"
    "math/rand"
)

func squaredNumber() int64 {
	var x int64
	fmt.Scanf("%d", &x)
	return rand.Int63n(x)
}
