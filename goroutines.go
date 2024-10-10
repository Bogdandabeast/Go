package main 

import "fmt"


func main() {

	fmt.Println("la funcion main, empezando....")

	go func() {

		fmt.Println("Hola desde Go Routine")

	}()

	time.Sleep(10 * time.Millisecond)

	fmt.Println("....Fin")



}