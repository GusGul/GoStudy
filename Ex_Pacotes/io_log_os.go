package main

import (
	"io"
	"log"
	"os"
)

func main() {
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
	if _, err := io.WriteString(file, "Olá, mundo!\n"); err != nil {
		log.Fatal(err)
	}
	// Ler o arquivo log.txt
	os.ReadFile("log.txt")
}
