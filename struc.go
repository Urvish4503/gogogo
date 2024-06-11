package main

type Cube struct {
	length, width, height float64
}

func (c *Cube) area() float64 {
	return c.length * c.width
}

func (c *Cube) volume() float64 {
	return c.length * c.width * c.height
}

type Person struct {
	Name string
	Age  int
	Job  bool
}

func makePerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}

	p.Job = false

	return &p
}

func personWithJob() *Person {
	p := Person{
		Name: "Uncle bob",
		Age:  45,
		Job:  true,
	}

	return &p
}
