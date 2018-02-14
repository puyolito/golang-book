/*package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println(number, "Fizz")
		} else if number%5 == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
*/
/*
package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3))

	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
*/
package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(fizzbuzz(number))
	}
}

func fizzbuzz(number int) (int, string) {

	if number%15 == 0 {
		return number, "FizzBuzz"
	
	} else if number%3 == 0 {
		return number, "Fizz"

	} else if number%5 == 0 {
		return number, "Buzz"

	} else {
		return number, ""
	}
}