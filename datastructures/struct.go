package main

import "fmt"

type Vertex struct { //Sctuct es una coleccion de campos
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} //Se instancia Vertex con valores x = 1 y y = 2
	v.X = 4           //Se asigna valor 4 a la variable x de Vertex de v
	fmt.Printf("Los valores son %v y %v", v.X, v.Y)
}
