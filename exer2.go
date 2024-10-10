package main 

import "fmt"

func main() {

	intArray := [...]int32{1,2,3}

	fmt.Println(intArray)

	cantidad := 12

	imprimirNumeros(cantidad)

}

func imprimirNumeros(cantidad int) {

	for i := 0; i < cantidad; i++ {

		fmt.Println(i)

	}

}