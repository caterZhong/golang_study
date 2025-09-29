package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("员工姓名:%s\n员工年龄:%d\n员工号:%s\n", e.name, e.age, e.EmployeeID)
}

func main() {
	e := Employee{
		Person: Person{
			name: "小明",
			age:  12,
		},
		EmployeeID: "123"}
	e.PrintInfo()

}
