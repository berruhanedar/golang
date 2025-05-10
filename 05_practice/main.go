/*
TASK 1 :
package main

import "fmt"

func main() {

	/* var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true */

/* var studentName  = "Johm Doe"
var grade  = 77
var isPassed = true */

/* studentName := "John Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}

*/

/*
TASK 2 :
package main

import "fmt"

func main() {

	/* var studentName, grade, isPassed = "John Doe", 77, true */
/* studentName, grade, isPassed := "John Doe", 77, true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

} */

/*
TASK 3 :
"Declaration , Assign , Initialization , Initial Value" kavramlarını açıklayınız
Declaration : Değişkeni tanımlıyorsun . Bu haliyle de çalışır çünkü zero values var
Assign : Oluşturduğun değişkene bir değer atıyorsun
Initialization : Direkt oluşturum değer atamak = var studentName string = "John Doe"
Initial Value : Default value


package main

import "fmt"

func main() {

	var studentName string = "John Doe"

	studentName = "Mahmut Erdem"

	fmt.Println(studentName)

} */

// TASK 4 : difference between := and = , double declaration

package main

import "fmt"

func main() {

	/* 	var studentName string = "John Doe"
	   	studentName = "Mahmut Erdem" */

	studentName := "John Doe"
	studentName = "Berru Hanedar "

	fmt.Println(studentName)

}
