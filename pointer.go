package main

func main() {
	// Ponteiro
	x := 10
	y := &x
	*y = 20
	println(x)
	println(y)
	println(*y)
	println(&x)
	println(&y)
}
