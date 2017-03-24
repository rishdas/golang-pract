package main

import "fmt"
import "bufio"
import "os"

func main() {
	countMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		countMap[input.Text()]++
	}
	for line, n := range countMap {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
