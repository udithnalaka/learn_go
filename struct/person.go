package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

// inner struct
type Engineer struct {
	person Person
	job string
}

func main() {
	p1 := Engineer{
		person: Person{
			firstName: "Udith",
			lastName:  "Nalaka",
			age:       33,
		},
		job: "Software Engineer",
	}

	p2 := Engineer{
		person: Person{
			firstName: "Iranga",
			lastName:  "Rana",
			age:       32,
		},
		job: "Civil Engineer",
	}

	fmt.Println(p1.person.firstName, p1.person.lastName, p1.person.age, p1.job)
	fmt.Println(p2.person.firstName, p2.person.lastName, p2.person.age, p2.job)
}
