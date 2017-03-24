package main

import "fmt"
import "os"

func main() {
	var args, sep string

	for  _, word := range os.Args[1:] {
		args += sep + word
		sep = " "
	}
	fmt.Println(args)
}
