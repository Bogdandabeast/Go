package main

import (
	"errors"
	"fmt"
)

// Funcion main que ejecuta el codigo
func main() {

	printValue := "Hola Programa"
	printMe(printValue)

	//Llamo a la funcion pasandole 0 y 2 como argumento, guardo el output en las variables
	var result, resto, err = intDivision(10, 2)

	//compruebo si el error es distinto de nil , quiere decir que ha habido error y ya no es nil el valor

	if err != nil {

		fmt.Println(err)

	} else {
		fmt.Printf("Formateo todo con el metodo printf de fmt y me sale asi: %v %v", result, resto)
	}

}

//func nombre(valor tipo) tipo que devuelve{}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// Se puede devolver mas de una cosa pero hay que especificar
func intDivision(num1 int, num2 int) (int, int, error) {

	var err error

	if num1 == 0 {
		err = errors.New("No se puede dividir por cero")

		return 0, 0, err //Devuelve dos numeros y el mensaje de error

	}

	result := num1 / num2
	resto := num1 % num2
	return result, resto, err //devuelve <nil>

}
