package main

import (
	"fmt"
	"os"
	"strconv"
)

func convertDecimalToBase(decimal int, base int) string {
	if base < 2 || base > 32 {
		return "La base debe estar entre 2 y 32"
	}

	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	result := ""

	for decimal > 0 {
		remainder := decimal % base
		result = string(digits[remainder]) + result
		decimal = decimal / base
	}

	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Se requieren dos argumentos: un número decimal y la base a la que se convertirá.")
		return
	}

	decimal, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("El primer argumento debe ser un número decimal válido.")
		return
	}

	base, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("El segundo argumento debe ser un número entre 1 y 32.")
		return
	}

	result := convertDecimalToBase(decimal, base)
	fmt.Printf("%d en base %d es %s\n", decimal, base, result)
}
