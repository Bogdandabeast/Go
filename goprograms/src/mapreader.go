package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	alreadyseen := make(map[string]bool)

	for in.Scan() {

		txt := in.Text()

		if _, found := alreadyseen[txt]; found {
			continue
		}

		alreadyseen[txt] = true

		fmt.Println(txt)

	}
	fmt.Printf("Los  valores repetidos: %v", alreadyseen)

}
