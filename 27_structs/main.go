/* package main

import "fmt"

func main() {

	var employee struct { // structure
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age)

	employee.name = "Berru"
	employee.age = 24
	employee.isMarried = true

	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)

	var employee2 struct { // structure
		name      string
		age       int
		isMarried bool
	}

	employee2.name = "Gul"
	employee2.age = 5
	employee2.isMarried = false

	fmt.Printf("%#v\n", employee2)
	fmt.Println(employee2.name)
	fmt.Println(employee2.age)
	fmt.Println(employee2.isMarried)

}
*/

/////////////////////////////////////////////////////////////////////////////////

/* package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
}

func main() {

	var e1 employee
	e1.name = "Berru"
	e1.age = 24
	e1.isMarried = true

	var e2 employee
	e2.name = "Gul"
	e2.age = 5
	e2.isMarried = false

	e3 := employee{
		name:      "Kemal",
		age:       3,
		isMarried: false,
	}

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
	fmt.Printf("%#v\n", e3)

}  */

////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
	kids      []string
}

func main() {

	e1 := employee{
		name:      "Berru",
		age:       24,
		isMarried: true,
		kids: []string{
			"Laisy",
			"Chris",
		},
	}

	//fmt.Printf("%#v\n", e1)
	fmt.Println(e1)
	fmt.Println(e1.kids)
	fmt.Println(e1.kids[0])

}
