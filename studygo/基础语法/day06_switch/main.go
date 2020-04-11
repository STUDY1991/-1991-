package main

import "fmt"

var dogSex = "M"

func main() {
	switch dogSex {
	case "M":
		fmt.Println("a Male dog")
	case "F":
		fmt.Println("a Female dog")
	}
}
