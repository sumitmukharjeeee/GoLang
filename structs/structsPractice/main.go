package main
import "fmt"
type Employee struct{
	Name string
	Role string
	Salary int
	isRemote bool
}
func main(){
	Employee1 := Employee{
		Name: "Alex Navalny",
		Role: "Principal Architect",
		Salary:1000000,
		isRemote: true,
	}
	fmt.Println(Employee1.Name)
	fmt.Println(Employee1.Role)
	Employee1.Salary = 300000
	fmt.Println(Employee1.Salary)
}