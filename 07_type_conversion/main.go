package main

import "fmt"

func main() {

	/* 	x := 10
	   	y := 10.0

	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)

	   	// fmt.Println(x + y) -> Farklı veri tiplerini toplayamazsın - type mismatch

	   	// Type Conversion -> type(value)
	   	fmt.Println(x + int(y)) */

	/* var x int8 = 10
	var y int16 = 10

	fmt.Println(x + y) // mismatched types hatası olur */

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y) */

	/* var x int16 = 128
	var y int8

	y = int8(x)

	fmt.Println(y) */

	/* x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(x + int(y)) // String int'e dönüşmez , çalışmaz */

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)

}
