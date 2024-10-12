package main

import "fmt"

type Vertex struct { //Struct Vertex 
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} //Declaro Vertex en v
	p := &v //Guarda la referencia en la memoria al v
	fmt.Println(&p) // & accede a la referencia , * accede al valor de la variable, 
	p.X = 1e9 //guarda el valor 1e9 en el p.X que lo que hace es guardarlo en v.X 
	fmt.Println(v)
}
