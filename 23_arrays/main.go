package main

import "fmt"

func main() {
	/* city1 := "istanbul"
	city2 := "roma"
	city3 := "tahran"
	city4 := "belgrad"

	/* fmt.Println(city1, city2, city3, city4) */

	// cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	// cities := [...]string{"istanbul", "roma", "tahran", "belgrad"}

	/* fmt.Println(cities)

	fmt.Println(cities)
	fmt.Println(cities[0]) // zero based index
	fmt.Println(cities[3])

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	var myArr [5]int
	fmt.Println(myArr) */

	// myArr[0] = "istanbul" - hata olur çübkü arrayler tek tip veri olur . Buradaki arrayin tipi int
	/* myArr[0] = 100
	fmt.Println(myArr)
	myArr[len(myArr)-1] = 200
	fmt.Println(myArr)
	*/
	/* var myArr3 [5]int
	var myArr4 [4]int

	if myArr3 == myArr4 { // eşit olmadıkları için program çalışmaz
		fmt.Println("MESAJ")
	}
	*/

	/* var myArr3 [5]int
	var myArr4 [4]int

	fmt.Println(myArr3)
	fmt.Println(myArr4) */

	/* cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	cities[0] = "ANKARA"
	fmt.Println()

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	} */

	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	myArr = mySquare(myArr) // First Class Functions

	fmt.Println(myArr)

	// FOR --- RANGE
	// for index, value := range myArr

	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	/* 	for index, city := range cities {
		fmt.Println(index, city)
	} */

	for _, city := range cities {
		fmt.Println(city)
	}

}

func mySquare(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}

	return arr

}
