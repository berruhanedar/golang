/* package main

import "fmt"

func main() {
	x := 75
	var y float64
	y = float64(x)

	fmt.Println(y)

}
*/

/* package main

import "fmt"

func main() {
	x := 5
	y := 10

	fmt.Println("X:", x, "Y:", y)

	x, y = y, x // x = y , y = x

	fmt.Println("X:", x, "Y:", y)

} */

/* // non english variable names
// UTF-8
package main

import "fmt"

func main() {

	yaş := 24
	名稱 := "Berru"

	fmt.Println("Name:", 名稱, "Age:", yaş)

}
*/

/* // Shadowing

package main

import "fmt"

func main() {
	x := 5

	if true {
		x := 10
		x++
		fmt.Println(x)
	}

	fmt.Println(x)

}
*/

// 5 : 40 as a string

package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 65

	s := string(x)

	fmt.Printf("%v, %T\n", x, x) // 65
	fmt.Printf("%v, %T\n", s, s) // A

	y := strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y) // "65"
}
