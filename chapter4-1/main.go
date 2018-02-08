package main

import "fmt"

func main() {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	yy := 12345
	fmt.Println(yy)
	fmt.Printf("Type: %T\n", yy)

	/*cannot assign to xx
	const xx string = "Hello, World"
	xx = "Other string"
	*/

	//Example
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//Example
	v1, v2 := "first", "sec"

	fmt.Println(v1)
	fmt.Println(v2)

	//Example Swap Position
	v3, v4 := "third", "fourth"
	v3, v4 = v4, v3
	
	fmt.Println(v3)
	fmt.Println(v4)

}

/* Error Case
func f() {
	fmt.Println(x)
}
*/