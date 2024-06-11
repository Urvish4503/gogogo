package main

import "fmt"

func sum(nums ...int) {
	ans := 0

	for i, num := range nums {
		fmt.Println(i, num)
		ans += num
	}

	fmt.Println(ans)
}
