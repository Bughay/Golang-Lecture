package main

import (
	"fmt"
)

func main() {
	struct_test()
}

func basic_types() {
	var integer int = 10
	var float float64 = 10.11
	var str string = "hi"
	var boolean bool = true

	fmt.Println("Basic data types")
	fmt.Println(integer, float, str, boolean)

	fmt.Println(&integer, &float, &str, &boolean)
}

func aggregate_types() {
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Aggregate Data Types")
	fmt.Println("array")

	fmt.Println(arr)
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
	fmt.Println(&arr[3])
	fmt.Println(&arr[4])

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice")
	a := &arr[0]
	fmt.Println(s)
	fmt.Println(a) //
}

type Employee struct {
	ID       int
	Name     string
	Salary   float64
	Active   bool
	Projects []string          // slice of strings
	Contact  map[string]string // map
	Manager  *Employee         // pointer to another Employee
}

func struct_test() {

	emp := Employee{
		ID:       101,
		Name:     "John Doe",
		Salary:   75000.50,
		Active:   true,
		Projects: []string{"Project A", "Project B"},
		Contact:  map[string]string{"email": "john@company.com", "phone": "123-456-7890"},
	}
	fmt.Println("Struct Employee no pointer")

	fmt.Println(emp)
	fmt.Println(emp.ID, emp.Name, emp.Salary)
	fmt.Println(&emp)
	fmt.Println(emp.Name)

	fmt.Println(&emp.Name)

	emp.Name = "mark"
	fmt.Println(emp.Name)

	fmt.Println(&emp.Name)

}
