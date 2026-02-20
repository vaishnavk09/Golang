package main

import "fmt"

func main() {

	i := 1
	//while is not available in golang but we can use for loop to achieve the same result
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// calssic for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)

	}
	// for range loop
	for i := range 12 {
		fmt.Println(i)
	}
}
