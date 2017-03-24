package main

import "fmt"
import "os"

func main() {
	var args, sep string
	
	for i := 1; i < len(os.Args); i++  {
		args += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(args)
}
