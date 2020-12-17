package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")

	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := s[1:3]
	fmt.Println("Range", l)

	m := s[:4]
	fmt.Println(m)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twod := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twod[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println("2D", twod)

}
