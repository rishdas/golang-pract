package main

import "fmt"

func main() {
	i := 0
	sum := 0
	for i = 0; i<10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	for sum < 1000 {
		sum += sum
	}
	if (sum > 1000) {
		fmt.Println(sum)
	} else {
		fmt.Println(sum - 1000)
	}
}
