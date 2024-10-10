package main 

import "fmt"

func main() {

	var intArr [3] int32 //instanci array de tamaño 3 de tipo de dato int32

	intArr2 := [3] int32{4,5,6}

	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3 

	

	fmt.Printf("los numeros de los arrays son $v $v", intArr , intArr2)

	fmt.Println(intArr[0])

	fmt.Println(intArr[1:3])


	


	



}

