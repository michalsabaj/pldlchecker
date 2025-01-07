package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/michalsabaj/pldlchecker/hasher"
)

type PersonalData struct {
	FirstName           string
	LastName            string
	DriverLicenseNumber string
}
type Settings struct {
	DriverLicenseNumberCharLimit int
	NameCharLimit                int
	Debug                        bool
	APIAddress                   string
}

func setSettings() Settings {
	return Settings{
		DriverLicenseNumberCharLimit: 8,
		NameCharLimit:                80,
		Debug:                        true,
		APIAddress:                   "https://moj.gov.pl/nforms/api/UprawnieniaKierowcow/2.0.10/data/driver-permissions?hashDanychWyszukiwania=",
	}
}

type ResponseData struct {
	//document confirming the permissions to drive
	DokumentPotwierdzajacyUprawnienia struct {
		//document type
		TypDokumentu struct {
			Kod     string `json:"kod"`
			Wartosc string `json:"wartosc"`
		} `json:"typDokumentu"`
		//document number
		SeriaNumerBlankietuDruku string `json:"seriaNumerBlankietuDruku"`
		//document issuer
		OrganWydajacyDokument struct {
			Kod     string `json:"kod"`
			Wartosc string `json:"wartosc"`
		} `json:"organWydajacyDokument"`
		//document expiration date
		DataWaznosci string `json:"dataWaznosci"`
		//document status
		StanDokumentu struct {
			StanDokumentu struct {
				Kod     string `json:"kod"`
				Wartosc string `json:"wartosc"`
			} `json:"stanDokumentu"`
			//reasons for the change of status
			PowodZmianyStanu []any `json:"powodZmianyStanu"`
		} `json:"stanDokumentu"`
		//categories of vehicles the driver is allowed to drive
		DaneUprawnieniaKategorii []struct {
			//category (like B)
			Kategoria string `json:"kategoria"`
			//expiration date of the category permission
			DataWaznosci string `json:"dataWaznosci"`
		} `json:"daneUprawnieniaKategorii"`
	} `json:"dokumentPotwierdzajacyUprawnienia"`
	// information from the system
	Komunikaty []any `json:"komunikaty"`
}

// unmarschal the json response
func unmarshalResponse(body []byte) ResponseData {
	var data ResponseData
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Unmarshalling failed, error: %s\n", err)
		os.Exit(1)
	}
	return data
}

func main() {
	//input first name, last name and driver license number separated by / and limit the input to 80 characters for first name and last name and 8 for driver license number
	fmt.Println(
		"Podaj imie, nazwisko i numer prawa jazdy.\nKazde z danych oddziel /.\nPrzyklad: Jan/Kowalski/A123456\nLimit znakow wynosi 80.\n-----")

	var pd PersonalData
	var input string
	fmt.Scanf("%s", &input)
	//separating the input
	pd, err := parseInput(input)
	if err != nil {
		fmt.Println("Błąd przy parse.", err)
		os.Exit(1)
	}
	settings := setSettings()
	//debugging
	if settings.Debug {
		fmt.Println("[DEBUG]: Imie: " + pd.FirstName)
		fmt.Println("[DEBUG]: Nazwisko: " + pd.LastName)
		fmt.Println("[DEBUG]: Prawo jazdy: " + pd.DriverLicenseNumber)
	}
	if len(pd.FirstName) > settings.NameCharLimit || len(pd.LastName) > settings.NameCharLimit {
		fmt.Printf("Za duzo znakow w imieniu lub nazwisku! (max. %d)\n", settings.NameCharLimit)
	}
	if len(pd.DriverLicenseNumber) > settings.DriverLicenseNumberCharLimit {
		fmt.Printf("Za duzo znakow w numerze prawa jazdy! (max. %d)\n", settings.DriverLicenseNumberCharLimit)
	}

	//hashing the input and storing it in a variable
	hash := hasher.HashDane(pd.FirstName, pd.LastName, pd.DriverLicenseNumber)
	if settings.Debug {
		fmt.Printf("[DEBUG]: Hash: %s\n", hash)
	}
	//sending the request and harvesting the response.
	res, err := http.Get(settings.APIAddress + hash)
	if err != nil {
		fmt.Printf("Request failed, error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response code: %d\n", res.StatusCode)
	//code 200 - ok, code 400 - not found

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Reading response failed, error: %s\n", err)
			os.Exit(1)
		}
		data := unmarshalResponse(body)

		fmt.Printf("stan dokumentu: %s\n", data.DokumentPotwierdzajacyUprawnienia.StanDokumentu)
		if data.DokumentPotwierdzajacyUprawnienia.StanDokumentu.StanDokumentu.Wartosc == "Wydany" {
			fmt.Println("Uprawnienia do kierowania pojazdami:")
			for _, kategoria := range data.DokumentPotwierdzajacyUprawnienia.DaneUprawnieniaKategorii {
				fmt.Printf("Kategoria: %s, data waznosci: %s\n", kategoria.Kategoria, kategoria.DataWaznosci)
			}
		} else {
			fmt.Println("Brak uprawnien do kierowania pojazdami.")
		}
		//status 400
	} else if res.StatusCode == http.StatusBadRequest {
		//Person with the given data does not exist in gov database
		fmt.Println("Osoba o podanych danych nie istnieje.")
	} else {
		fmt.Println("blad")
	}
	defer res.Body.Close()
}

func parseInput(input string) (PersonalData, error) {
	parts := strings.Split(input, "/")
	if len(parts) != 3 {
		//incorrect input format
		return PersonalData{}, fmt.Errorf("niepoprawny format wejściowy")
	}
	return PersonalData{
		FirstName:           parts[0],
		LastName:            parts[1],
		DriverLicenseNumber: parts[2],
	}, nil
}
