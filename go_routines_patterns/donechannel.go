package main

import "fmt"
import "time"

func main() {

	//canal done.

	done := make(chan bool)

	go doWork(done)

	//Cuando el hilo main acaba usa el canal main para acabar con los hijos

	time.Sleep(time.Second * 3)

}

//la funcion recibe el canal done

func doWork(done <-chan bool) {
	for {
		select {
		//el caso que hace que la funcion termine
		case <-done:
			return
		default:
			fmt.Println("Haciendo trabajo")
		}
	}
}
