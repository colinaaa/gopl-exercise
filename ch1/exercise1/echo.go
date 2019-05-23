// Exercise 1.1:Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args {
		fmt.Print(arg, " ")
	}
	fmt.Println("")
}
