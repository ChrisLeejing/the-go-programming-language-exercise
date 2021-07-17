package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

var dilbert Employee

// func EmployeeById(id int) Employee {
// 	return Employee{ID: 1}
// }

func EmployeeById(id int) *Employee {
	return nil
}

func main() {
	fmt.Println(EmployeeById(dilbert.ManagerID).Position)
	id := dilbert.ID
	// byId := EmployeeById(id)
	// byId.Salary = 0
	// EmployeeById(id).Salary = 0
	EmployeeById(id).Salary = 0
}