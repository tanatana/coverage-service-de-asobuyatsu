package main

import (
	"fmt"
	"github.com/tanatana/coverage-service-de-asobuyatsu/add"
)

func main() {
	fmt.Printf("%#d\n", getSum(1,1))
	fmt.Printf("%#d\n", getHikizan(2,1))
}

func getSum(a, b int) int {
	return add.Add(a, b)
}

// pr でつっこみもらうメソッド
func getHikizan(a, b int) int {
	return a - b
}