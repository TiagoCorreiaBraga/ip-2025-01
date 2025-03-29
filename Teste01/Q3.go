package main

import (
	"fmt"
)

func main() {
	soma := 0

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		soma += i
	}

	fmt.Printf("Soma total: %d\n", soma)
}
