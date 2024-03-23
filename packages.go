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

	// Pacote io, log e os - manipulação de arquivos
	// Pacote io
	if _, err := io.WriteString(os.Stdout, "Olá, mundo!\n"); err != nil {
		log.Fatal(err)
	}
	// Pacote os
	// Escrever no arquivo log.txt criado com io
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "Olá, mundo!"); err != nil {
		log.Fatal(err)
	}
	// Ler o arquivo log.txt
	os.ReadFile("log.txt")
}
