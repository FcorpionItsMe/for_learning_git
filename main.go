package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Input A:")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Input B:")
	fmt.Scanf("%d\n", &b)
	fmt.Printf("Diference of a and b = %d", a-b);
}
