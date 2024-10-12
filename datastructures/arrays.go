package main

import (
	"fmt"
)

func main() {

	fmt.Println("WELCOME TO ARRAYS")

	//Arrays son fijos y estaticos.
	//var nombre [tamaño] tipo

	var intARR [3]int32

	//asignar valor

	intARR[0] = 240

	//declarar e iniciar la array

	var intARR2 [3]int32 = [3]int32{1, 2, 3}

	//otra manera de iniciar la array. [...] Quiere decir que el tamaño lo defina el compilador que en este caso sera 6

	intARR3 := [...]int32{1, 2, 3, 4, 5, 6}

	fmt.Println(intARR)

	fmt.Println(intARR2)

	fmt.Println(intARR3)

	//Slices son envoltorios para arrays pero con mas funcionalidades.

	intSlice := [3]int32{4, 5, 6}

	intSlice = append(intSlice, 7)

}
