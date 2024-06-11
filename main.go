package main

import "fmt"

func main() {
	//sliceFun()
	//mapThing()
	//sum(1, 2, 3, 5, 67, 8)
	//nextInt := intSeq()
	//
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//
	//newInts := intSeq()
	//fmt.Println(newInts())
	//fmt.Println(newInts())
	//fmt.Println(nextInt())

	ten := 0

	pointer(&ten)
	fmt.Println(ten)

	var c = makePerson("baba", 21)

	fmt.Println(c)

	cc := personWithJob()
	fmt.Println(cc.Job)
	cc.Job = false

	fmt.Println(cc.Job)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
