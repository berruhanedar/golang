package main

import "fmt"

var packVar = "Package Scope"

// packVar := "Package Scope" - kısa gösterim paket değişkenlerinde kullanamazsın
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar) // it works
	}

	// it doesn't work fmt.Println(blockVar)

	var funcVar = "Func Scope" // Buna öncelik verir bunu basar , paket düzeyindeki değil

	fmt.Println(funcVar)
	fmt.Println(packVar)

	myFunc()

	var name = "Berru"
	// var name := "Burcu" - Çalışmaz var olan variable!a tekrar bir şey atayamazsın
	name, surname := "Burcu", "Hanedar"
	fmt.Println(name, surname) // Burada çalışırım her iki işi birden yapar . Hem yeni değişken oluşturur hem de atama yapar
}

func myFunc() {
	fmt.Println(packVar)
	// fmt.Println(funcVar) - it does not work main içindeki
	fmt.Println(funcVar) // Paket düzeyindeki için
}
