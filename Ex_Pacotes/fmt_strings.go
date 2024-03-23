package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Pacote strings, formatações de strings
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.Contains("Hello, World!", "World"))
	fmt.Println(strings.Contains("Hello, World!", "Mundo"))
}
