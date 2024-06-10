package main

import (
	"fmt"
)

func sliceFun() {
	var s []string

	fmt.Println("len:", len(s), ", cap:", cap(s))
	fmt.Println("s == nil", s == nil)

	s = make([]string, 3)

	fmt.Println("len:", len(s), ", cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s =", s)
	fmt.Println("s == nil", s == nil)

	s = append(s, "c")

	c := make([]string, len(s))
	copy(c, s)

	l := c[2:4]

	fmt.Println("l =", l, "cap:", cap(c))

}
