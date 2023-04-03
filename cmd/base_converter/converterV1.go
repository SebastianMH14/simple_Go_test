package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Se necesitan al menos dos parámetros.")
		return
	}

	// Obtener los parámetros desde la consola
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: El primer parámetro debe ser un número decimal válido.")
		return
	}

	base, err := strconv.Atoi(os.Args[2])
	if err != nil || base < 2 || base > 32 {
		fmt.Println("Error: El segundo parámetro debe ser un número entre 2 y 32.")
		return
	}

	// Convertir el número a la nueva base
	result := strconv.FormatUint(uint64(num), base)

	// Imprimir el resultado
	fmt.Printf("%d escrito en base 10 en base %d es %s.\n", num, base, result)
}
