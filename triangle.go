package main

import "fmt"

var (
	base, heigth float64
)

func main() {
	fmt.Print("Enter base: ")
	//var base, heigth float64
	fmt.Scanf("%f", &base)
	fmt.Print("Enter heigth: ")
	fmt.Scanf("%f", &heigth)

	area := (base * heigth) / 2
	fmt.Println("the triangle area is:", area)

}
