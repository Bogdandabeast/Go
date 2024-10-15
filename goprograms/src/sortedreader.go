package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	var previo string

	for in.Scan() {
		txt := in.Text()

		if txt == previo {
			continue
		}

		/* if txt < previo {
			panic("archivo no esta ordenado")
		} */

		previo = txt

		fmt.Println(txt)
	}

}
