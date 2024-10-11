package main

import "fmt"
import "time"

func  main()  {

	//Ejecucion async go. Se hace un fork del main y se crearan 3 otros main llamados children.

	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	//Pongo a dormir el hilo main para que a los hijos les de tiempo a acabar

	time.Sleep(time.Second * 2) 

	

	fmt.Println("hola mundo")

	
}

func someFunc(num string) {

	fmt.Println(num)

}