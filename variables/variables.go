package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//var name type

	//unsigned  variables guardan valores positivos solo

	var positivo uint = 45

	var variable int = 3

	var variablefloat float32 = 3.1342342

	//Se puede omitir el tipo de variable. En este caso se usara el tipo de dato almacenado

	var cadena = "mis"

	//Se puede omitir la palabra reservada var poniendo :

	cadena2 := "amigo son"

	//Constante, inmutable

	const Myconst = "Constante"

	//para declarar varias variables

	var1, var2, var3 := " Ronald", " Bogdan", " Pavel"

	fmt.Println(cadena + " " + cadena2 + var1 + var2 + var3 + Myconst)

	//No se pueden realizar operaciones con tipos distintos, hay que castear.

	var result float32 = variablefloat + float32(positivo)

	//Declaro string

	var myString string = "\nHello World"

	//Imprimo el tamaño de los bits que ocupa este caracter en utf8

	fmt.Println(utf8.RuneCountInString("ñ"))

	//Creo una variable de tipo rune

	var myRune rune = 'c'

	//Imprimo su orden en utf8

	fmt.Println(myRune)

	//Declaro un boolean false

	var myBoolean bool = false

	fmt.Println(myBoolean)

	fmt.Printf("Mis primeras variables: %v %v %v", variable, positivo, result)
	fmt.Println(myString)
}
