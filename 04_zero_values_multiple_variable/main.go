package main

import "fmt"

func main() {

	// var name = "Berru"
	// var age uint8 = 24
	// var isStudent = true
	// var weight float32 = 58.5

	/* var (
		name      string = "Berru"
		age       uint8  = 24
		isStudent bool   = true

		height int     = 158
		weight float32 = 58.5
	) */

	/* var name, age, isStudent, weight, height = "Berru", 24, true, 58.5, 158 */

	/* name, age, isStudent, weight, height := "Berru", 24, true, 58.5, 158 */

	var name string
	var age int
	var weight float32
	var isStudent bool

	fmt.Println(name)      // Zero Values = string --> ""
	fmt.Println(age)       // Zero Values = 0
	fmt.Println(isStudent) // Zero Values = false
	fmt.Println(weight)    // Zero Values = 0

}
