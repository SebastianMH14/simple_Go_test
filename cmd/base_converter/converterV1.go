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
	numStr := os.Args[1]
	baseOrigen, err := strconv.Atoi(os.Args[2])
	if err != nil || baseOrigen < 2 || baseOrigen > 32 {
		fmt.Println("Error: El segundo parámetro debe ser un número entre 2 y 32.")
		return
	}

	// Convertir el número a decimal
	num, err := strconv.ParseUint(numStr, baseOrigen, 64)
	if err != nil {
		fmt.Printf("Error: No se pudo convertir '%s' a base %d.\n", numStr, baseOrigen)
		return
	}

	// Obtener la base de destino
	baseDestino := 10
	if len(os.Args) >= 4 {
		baseDestino, err = strconv.Atoi(os.Args[3])
		if err != nil || baseDestino < 2 || baseDestino > 32 {
			fmt.Println("Error: El tercer parámetro debe ser un número entre 2 y 32.")
			return
		}
	}

	// Convertir el número a la nueva base
	result := strconv.FormatUint(num, baseDestino)

	// Imprimir el resultado
	fmt.Printf("%s en base %d es %s en base %d.\n", numStr, baseOrigen, result, baseDestino)
}
