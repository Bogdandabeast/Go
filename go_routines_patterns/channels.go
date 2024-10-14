package main

import "fmt"

func main() {

	//Un channel de Go es como una pila de fifo, first out first in que le da informacion a las funciones asincronas y tambien al main

	myChannel := make(chan string)

	anotherChannel := make(chan string)

	//funcion anonima que se crea y se llama en el momento. con go se hace un fork al main. El main espera a que la informacion se introduzca al canal

	go func() {

		myChannel <- "datos para el primer canal"
	}()

	go func() {

		anotherChannel <- "datos para el segundo canal"
	}()

	//fork child se reune con el main

	/* msg := <- myChannel */

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)

	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

	/* fmt.Println(msg) */

}
