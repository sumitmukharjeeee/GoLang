package main
// Every go file belongs to a package

//For an executable program we use package main generally

// FMT package imported


import "fmt"

func main(){
	// name := "Sumit" //String
	// age := 22 //Integer
	// prince:=15.45 //float64
	// active := true //Boolean

	//Go is statically typed labguage, so we need to specify the type

	var name string = "Sumit"
	var age int = 22

	//well useually you see name:= "Sumit" and age := 22, but we can also use var name string = "Sumit" and var age int = 22
	// becuase go can infer types
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)

	// Go is strongly typed
	// This is where our python intuitions need to cahgne

	money:=2200
	// money is int
	//i cant do moneu = "two thousand" because money is int and "two thousand" is string
	fmt.Println("I have", money, "dollars")

	// functions
	result := add(5, 10)
	fmt.Println("The sum is:", result)

	subResult := subtract(10,5)
	fmt.Println("The sub is:", subResult)

	subDivide, err := divide(15.5, 5.5)
	if err != nil{
			fmt.Println(err)
			return
	}

	fmt.Println(subDivide)
	// Going to be extremely familiar when you write go

	// Go doesnt use exceptions like python

	// No paranthesis for for loop
	var ages int = 21
	if ages>=20{
		fmt.Println("Adult")
	}else{
		fmt.Println("Not adult")
	}

	// Loops

	for i:=0;i<=5;i++{
		fmt.Println(i)
	}

}

func add(a int, b int) int {
	return a + b
}

func subtract(a,b int) int{
	return a-b
	// can shorten the parameter declaration when consecutive parameters have same types
}


	// Multiple return values
	func divide(a, b float64)(float64, error){
		if b ==0{
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		return a/b, nil
	}
