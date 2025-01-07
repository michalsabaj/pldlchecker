package data

//struct for personal data and others
type PersonalData struct {
	FirstName           string
	LastName            string
	DriverLicenseNumber string
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
