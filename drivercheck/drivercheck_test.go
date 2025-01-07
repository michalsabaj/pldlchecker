package drivercheck_test

import (
	"testing"

	"github.com/michalsabaj/pldlchecker/drivercheck"
)

func TestIsDriverLicenseValid(t *testing.T) {
	got, err := drivercheck.IsDriverLicenseValid("Andrzej", "Kowalski", "A123456")

	want := false

	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsDriverLicenseValidEmptyFields(t *testing.T) {
	got, err := drivercheck.IsDriverLicenseValid("", "", "")

	want := false

	if err == nil {
		t.Errorf("got %v, want %v", err, "one or more fields are empty")
	}

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
func TestGetDriverLicenseHash(t *testing.T) {
	got, err := drivercheck.GetDriverLicenseHash("Oliwier", "Społeczny", "A123456C")

	want := "1C11FAABFB21ACD3CACB6152E0AA6BC9"

	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
