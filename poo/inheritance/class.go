package main

import "fmt"

type Person struct {
	name string
	age  int
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d", p.name, p.age)
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v", ftEmployee)
}
