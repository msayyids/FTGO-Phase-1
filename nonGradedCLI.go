package main

import "fmt"

func NonGradedCLI() {
	var nama string
	fmt.Println("masukkan nama:")
	fmt.Scan(&nama)
	fmt.Printf("hello %s", nama)
}
