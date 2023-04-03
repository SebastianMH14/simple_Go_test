package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Se necesitan tres argumentos: dos números y una operación (+, -, *, /)")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Los dos primeros argumentos deben ser números válidos.")
		return
	}
	oper := os.Args[2]
	var res float64
	switch oper {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Printf("no se puede dividir por cero")
			return
		}
		res = num1 / num2
	default:
		fmt.Printf("la operación debe ser +, -, * o /")
		return
	}

	sumStr := strconv.FormatFloat(res, 'f', 2, 64)
	if sumStr[len(sumStr)-3:] == ".00" {
		fmt.Printf("%.0f %s %.0f = %.0f\n", num1, oper, num2, res)
	} else {
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, oper, num2, res)
	}

}
