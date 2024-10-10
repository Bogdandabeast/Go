package main 

import "fmt"
		"github.com/spf13/cobra

func main() {

	fmt.Println("mi todoapp")

	if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }



}


var rootCmd = &cobra.Command{
    Use:   "ronald",
    Short: "Mi aplicación CLI",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("¡Hola desde mi primera app que se ejecuta con ronald!")
    },
}