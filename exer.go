package main

import "fmt"

func main() {

	var intNum int = 323222

	var numero int = 2

	intNum = intNum + 1

	fmt.Println(intNum)



	fmt.Println(sumar(numero))

	



	 

}

func sumar(numero int) int {
	return numero + numero
}