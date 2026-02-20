package main

import "fmt"

func main() {
	var name string = "golang"
	var age int = 21
	var iscool = true // if we dont give datatype it will automatically assign the datatype based on the value

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(iscool)

	//short hand syntax
	line := "golang has short hand syntax"

	fmt.Println(line)

	const pi = 3.14
	fmt.Println(pi)

	fmt.Println("enter the radius:")
	var radius float64
	fmt.Scanf("%f", &radius)
	fmt.Println(radius)

	area := pi * radius * radius
	fmt.Println(area)

}
