package main

import "fmt"

func main() {

	//Si defines el tamaño del canal entonces es buffered y sino entonces es unbuffered. Para meter y sacar informacion al canal hay que usar funciones Go. Si es buffered la funcion asincrona no se bloqueara y funcionará como tal

	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	//cerramos el canal

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}
