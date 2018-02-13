package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Print("Enter Number: ")
		var number float64
		fmt.Scanf("%f\n", &number)

		if number == 22 {
			fmt.Println(number, "เจอแล้ว")
		} else if number > 50 {
			fmt.Println(number, "มากไป")
		} else {
			fmt.Println(number, "น้อยไป")
		}
	}
	fmt.Println("เกินพอ")
}