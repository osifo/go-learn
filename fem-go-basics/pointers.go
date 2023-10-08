package main

import "fmt"

func markAnniversary(weddingAge *int) {
	fmt.Println("Happy Wedding Anniversary!")

	*weddingAge++
}

func doWedding() (string, string, int) {
	var husband, wife = "Osifo", "Sonia"
	weddingAge := 0

	return husband, wife, weddingAge
}

func UpdateWeddingDetails() {
	husband, wife, weddingAge := doWedding()
	fmt.Printf("Happy married life to! %s and %s \n", husband, wife)

	markAnniversary(&weddingAge)
	markAnniversary(&weddingAge)
	markAnniversary(&weddingAge)

	fmt.Printf("You've now been married for %d years", weddingAge)
}
