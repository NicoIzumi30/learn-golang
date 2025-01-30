package main

// var Data int => ini public
// var data int => ini private

import "fmt"

func greet2(name string) {
	fmt.Printf("Hello %s\n", name)
}

func greet(name string) (text string) {
	// text = "Hello " + name
	// return
	return "Hello " + name //Bisa pakai ini atau atas nya
}

func add(x, y int) int {
	return x + y
}

// func addSub(x, y int) (int, int) {
// 	return x + y, x - y //langsung mereturn response
// }

func addSub(x, y int) (addRes int, subRes int) {
	addRes = x + y
	subRes = x - y
	return //deklarasikan dulu kedalam variable
}

func incremenet(var1 int) {
	var1 = var1 + 1 //tanpa pointer jadi tidak bisa
}

func incremenetPtr(var1 *int) {
	*var1 = *var1 + 5 //menggunakan pointer dengan menambahkan *
}

func main() {
	// fmt.Println(greet("World"))
	// greet2("Dunia")
	fmt.Println(add(3, 5))
	fmt.Println(addSub(3, 5))
	var1 := 1
	incremenet(var1)
	fmt.Println(var1)
	incremenetPtr(&var1)
	fmt.Println(var1)

	// cek address variable
	fmt.Printf("My address is %X, value is %d\n", &var1, var1)

	var var2 *int
	fmt.Printf("My value is %v\n", var2)

	var2 = new(int)
	*var2 = 5
	fmt.Printf("My value is %v\n", *var2)

	var2 = &var1
	*var2 = 100
	fmt.Printf("My value is %v\n", var1)
}
