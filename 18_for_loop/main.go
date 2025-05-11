package main

import "fmt"

func main() {

	// for <init statement>; <condition>; <post statement> { ----- }

	/* for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	}  */

	/* i := 0

	for ; i <= 10; i += 5 {
		fmt.Println(i)
	}

	fmt.Println(i) */

	/* for {  // Infinite Loop
		fmt.Println("Benim Adım Berru")
	}  */

	/* 	for i := 0; true; i += 5 {
		fmt.Println(i)
	} */

	/* 	for i := 0; false; i += 5 {
		fmt.Println(i)
	} */

	/* 	i := 10

	   	for i >= 0 {
	   		fmt.Println(i)
	   		i--
	   	} */

	/* 	for i := 0; i <= 10; i++ { // continue --> döngünün başına git

		if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	} */

	for i := 0; i <= 10; i++ {

		if i == 3 {
			break
		}

		fmt.Println(i)
	} // break --> döngüden çık

}
