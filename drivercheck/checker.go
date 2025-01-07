package drivercheck

import (
	"fmt"
	"log"

	"github.com/michalsabaj/pldlchecker/config"
	"github.com/michalsabaj/pldlchecker/handler"
	"github.com/michalsabaj/pldlchecker/hasher"
)

// IsDriverLicenseValid checks if the driver's license is valid.
// Requires the first name, last name, and driver's license number.
func IsDriverLicenseValid(firstName, lastName, driverLicenseNumber string) (bool, error) {
	cfg := config.GetConfig()
	//checking if the input data is correct
	if len(firstName) == 0 || len(lastName) == 0 || len(driverLicenseNumber) == 0 {
		return false, fmt.Errorf("one or more fields are empty")
	}
	if len(firstName) > cfg.NameCharLimit || len(lastName) > cfg.NameCharLimit || len(driverLicenseNumber) > cfg.DriverLicenseNumberCharLimit {
		return false, fmt.Errorf("input exceeds character limit")
	}
	//hashing the input data
	hash := hasher.HashDane(firstName, lastName, driverLicenseNumber)
	//sending the request and harvesting the response
	res, err := handler.HandleRequest(cfg.APIAddress + hash)
	if err != nil {
		return false, fmt.Errorf("failed to handle request: %w", err)
	}
	//json unmarshalling
	data, err := handler.UnmarshalResponse(res)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	//checking if the driver's license is valid
	isValid := data.DokumentPotwierdzajacyUprawnienia.StanDokumentu.StanDokumentu.Wartosc == "Wydany"
	if isValid {
		log.Printf("Driver's license status: %s", data.DokumentPotwierdzajacyUprawnienia.StanDokumentu)
	}
	fmt.Println("hi")
	return isValid, nil
}
