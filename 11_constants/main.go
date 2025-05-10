package main

import (
	"fmt"
	/* 	"math" */)

func main() {

	/* 5
	   3.14
	   "passed"S
	   true */

	/* r := 3.0 */

	/* areaOfCircle := 3.14 * (math.Pow(r, 2)) */

	/* const pi float64 = 3.14

	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle) */

	/* const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)
	*/

	// const x = 2
	// x = 5 constanta değer atayamayız
	// x++
	// x = x + 1

	// const ---> compile time
	// var ---> run time

	// var x = math.Pow(3, 4)
	// const x = math.Pow(3, 4)

	/* 	const x = 5

	   	fmt.Printf("%T, %v\n", x, x) */

	/* 	y := 3

	   	const x = y

	   	fmt.Printf("%T, %v\n", y, y)
		fmt.Printf("%T, %v\n", x, x) */

	/* 	const x = 1
	   	var y = 3

	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", x+y, x+y)
	*/

	const x, y = 3, 5

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

}
