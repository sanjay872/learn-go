package examples

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func Struct() {
	fmt.Println(Person{"sanjay", 25})
	fmt.Println(Person{name: "Sanjay", age: 25})
	fmt.Println(Person{name: "Sanjay"}) // default value is added for mission data
	fmt.Println(Person{age: 25})

	fmt.Println(NewPerson("Sanjay", 25))

	s := NewPerson("sanjay", 25)

	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(s.name)
	fmt.Println(&s.name)

	p := &s

	fmt.Println(p)
	fmt.Println((**p).name)
	fmt.Println(&((**p).name))

	q := struct {
		name string
		age  int
	}{
		"Sanjay",
		25,
	}

	fmt.Println(q)
}
