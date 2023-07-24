package main

import "fmt"

func NonGradedArraySlice1() {
	people := []string{"walt", "jessey", "skyler", "paul"}
	people = append(people, "hank")
	people = append(people, "marry")

	fmt.Println(people)

}
