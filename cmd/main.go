package main

import (
	"fmt"

	"github.com/michalsabaj/pldlchecker/drivercheck"
)

func main() {
	check, err := drivercheck.IsDriverLicenseValid("Andrzej", "Kowalski", "A123456")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Check: ", check)
}
