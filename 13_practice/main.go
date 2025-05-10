/* package main

import "fmt"

func main() {

	// TASK 1
	/* x := 50

	// x = x - 10  assignment statement

	x -= 10 // assigment operation

	fmt.Printf("%v , %T\n", x, x)

	x += 10

	fmt.Printf("%v , %T\n", x, x)

	x *= 10

	fmt.Printf("%v , %T\n", x, x)

	x /= 10

	fmt.Printf("%v , %T\n", x, x) */

// TASK 2
/*
	F := -40
	K := float64((F-32))/1.8 + 273

	fmt.Println(K) */

/*

age := 40

const myAge = age

fmt.Printf("%v, %T\n", myAge, myAge)

*/

/* package main

import "fmt"

func main() {

	const myAge = 40

	fmt.Printf("%v, %T\n", myAge, myAge)

} */

/* package main

import "fmt"

var x = 14

func main() {
	var x = 24
	fmt.Printf("%v , %T\n", x, x)
}
*/

//TASK 5 - const x = 4, const y = 5.4,  x + y?

/* package main

import "fmt"

const x = 14

func main() {

	const x = 4   // typeless
	const y = 5.4 // typeless

	fmt.Printf("%T, %v\n", x+y, x+y) // x in veri tipi float64

	fmt.Printf("%T, %v\n", x, x) // x in veri tipi int

	fmt.Printf("%T, %v\n", y, y)

} */

// 6-) const x float64 = 6.4 , y := 4 + x, y?

package main

import "fmt"

func main() {

	const x float64 = 6.4

	y := 4 + x

	fmt.Printf("%T, %v\n", y, y)

}
