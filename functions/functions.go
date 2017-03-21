package main

import "fmt"
import "math/cmplx"

var c, python, java bool

func add (x int, y int) int {
	return x + y
}
func divide (divr int, divd int) (quot int, rem int) {
	quot = divd / divr
	rem = divd % divr
	useless := 1
	useless = useless + 16
	fmt.Println(useless)
	return
}
func swap (a string, b string) (string, string) {
	return b, a
}
func main () {
	var i int
	var c_128 complex128
	const t_or_f  = true

	u_64 := 1
	c_128 = cmplx.Sqrt(-5 + 12i)
	
	fmt.Println("Hello World")
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(add(5, 2))
	fmt.Println(divide(5, 2))
	fmt.Println(i, c, python, java)
	//Printf
	fmt.Printf("Type : %T Value: %v\n", i, i)
	fmt.Printf("Type : %T Value: %v\n", u_64, u_64)
	fmt.Printf("Type : %T Value: %v\n", c_128, c_128)
	fmt.Printf("Type : %T Value: %v\n", t_or_f, t_or_f)
	
	
}
