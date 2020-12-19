package main

import (
	"fmt"
	"./sun/sun"
)

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	newSun := sun(2, 3)
	fmt.Println(newSun)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
