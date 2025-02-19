package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		fmt.Print("Write ", i, " as ")
		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}
	}

}
