package main

import "fmt"

func main() {
	i, j := 41, 2701

	p := &i         // Puntero a i
	fmt.Println(*p) // lee i a traves del puntero
	*p = 21         // Escribe en i a traves del puntero
	fmt.Println(i)  // muestra el nuevo valor de la i

	p = &j         // puntero a j
	*p = *p / 37   // divide j a traves del puntero
	fmt.Println(j) // muestra el nuevo valor de la j
}
