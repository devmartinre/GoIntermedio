package main

import "fmt"

type Person struct {
	name string
	age  int
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d", p.name, p.age)
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Time Employee"
}

type PrintInfo interface {
	getMessage() string
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
