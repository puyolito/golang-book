package main

import "fmt"
/*
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
*/
func main() {
	fmt.Print("Enter Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output :=  (input - 32) / 1.8
	fmt.Printf("%.2f\n", output)
}