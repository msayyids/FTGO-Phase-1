package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func soal1() {
	var nama string
	fmt.Println("masukkan nama : ")
	fmt.Scan(&nama)
	angka := rand.Intn(100)
	fmt.Printf("angka yang anda dapatkan : %d \n", angka)

	switch {
	case angka > 80:
		fmt.Printf("\nselamat %s anda sangat beruntung", nama)
	case angka > 60 && angka <= 80:
		fmt.Printf("\nselamat %s anda beruntung", nama)
	case angka <= 60 && angka > 40:
		fmt.Printf("\nmohon maaf %s anda kurang beruntung", nama)
	default:
		fmt.Printf("\nmohon maaf %s anda sial", nama)
	}
}

func soal2() {
	var nama string
	fmt.Println("\nmasukkan nama: ")
	fmt.Scan(&nama)

	var umurs string

	fmt.Println("Masukkan umur:")
	fmt.Scan(&umurs)

	umur, err := strconv.Atoi(umurs)
	if err != nil {
		fmt.Println("\nUmur harus berupa angka")

	} else {
		if nama != "" && umur > 0 && umur < 100 {
			if umur > 18 {
				fmt.Println("\nSilahkan masuk")
			} else {
				fmt.Println("\nDilarang masuk")
			}
		} else {
			fmt.Println("\nnama kosong atau umur yang dimasukkan tidak masuk akal")
		}
	}

}

func main() {
	soal1()
	soal2()
}
