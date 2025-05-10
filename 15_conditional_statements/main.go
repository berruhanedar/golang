package main

import "fmt"

func main() {

	/* // if <boolean expression> { code }

	x := 27

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	}

	if x%2 != 0 {
		fmt.Println(x, "tek sayıdır")
	} */

	if !false {
		fmt.Println("Mesaj gösterilecek")
	}

	if 5 > 3 {
		fmt.Println("Mesaj Gosterilecek mi?")
	}

	// if <boolean expression> {
	// 		code
	// } else {
	// 		code
	// }

	/* x := 4

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	} */

	// if <boolean expression> { code } else if <boolean expression> else { code }

	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}
}
