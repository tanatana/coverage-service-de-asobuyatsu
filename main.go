package main

import (
	"fmt"
	"github.com/tanatana/coverage-service-de-asobuyatsu/add"
)

func main() {
	fmt.Printf("%#d", getSum(1,1))
}

func getSum(a, b int) int {
	return add.Add(a, b)
}