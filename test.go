package main

import "fmt"

func Test() {
	var inputNama, alamat string

	fmt.Println("input nama:")
	fmt.Scanln(&inputNama)
	fmt.Println("input alamat :")
	fmt.Scanln(&alamat)
	fmt.Printf("namanya adalah %s, tinggalnya di %s", inputNama, alamat)

}
