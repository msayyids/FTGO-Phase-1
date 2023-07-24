package main

import "fmt"

func NonGradedArraySlice2() {
	identitas := []map[string]string{}

	identitas = append(identitas, map[string]string{"name": "hank", "gender": "M"})
	identitas = append(identitas, map[string]string{"name": "heisenverg", "gender": "M"})
	identitas = append(identitas, map[string]string{"name": "Skyler", "gender": "F"})

	for _, output := range identitas {
		namaPenduduk := output["name"]
		genderPenduduk := output["gender"]
		fmt.Printf("\nnama : %s", namaPenduduk)
		fmt.Printf("\ngender : %s", genderPenduduk)
	}

	fmt.Println("\n----------------------")

	for _, hasil := range identitas {
		nama := hasil["name"]
		gender := hasil["gender"]

		if nama != "" && gender == "M" {
			nama = "Mr " + hasil["name"]
		} else {
			nama = "Mrs " + hasil["name"]
		}

		fmt.Printf("\nnama : %s", nama)
		fmt.Printf("\ngender : %s", gender)

	}

}
