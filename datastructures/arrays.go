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

	//Slices son referencias a las arrays pero no guardan la informacion.

	slice := intARR3[0:3] //He creado un slice con los valores del 0 al 3  (1 2 3)

	fmt.Println(slice) //1 2 3

	//Ejemplo de declaracion de un Map

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Ronald": 16, "Frank": 15}
	fmt.Println(myMap2["Ronald"]) //16
	fmt.Println(myMap2)

}
