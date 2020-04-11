package main

import "fmt"

var dogAge = 10

func main() {
	if dogAge > 10 {
		fmt.Printf("A big Dog")
	} else if dogAge > 1 && dogAge <= 10 {
		fmt.Println("A small dog")
	} else {
		fmt.Println("A baby dog")
	}
}
