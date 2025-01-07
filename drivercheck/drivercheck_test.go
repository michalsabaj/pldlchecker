package drivercheck_test

import (
	"testing"

	"github.com/michalsabaj/pldlchecker/drivercheck"
)

func TestIsDriverLicenseValid(t *testing.T) {
	drivercheck.IsDriverLicenseValid("Andrzej", "Kowalski", "A123456")

}
