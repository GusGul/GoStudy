package main

import "fmt"

func main() {
	// A função recover() é utilizada para recuperar o controle do programa após um pânico.
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico! Deu erro amigão!")
}
