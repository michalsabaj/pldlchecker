package handler

//handle http request
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/michalsabaj/pldlchecker/data"
)

// handle http request
func HandleRequest(url string) ([]byte, error) {
	// Send request
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %w", err)
	}

	return body, nil
}

// handle json unmarshalling
func UnmarshalResponse(body []byte) (data.ResponseData, error) {
	var responseData data.ResponseData
	err := json.Unmarshal(body, &responseData)
	if err != nil {
		return data.ResponseData{}, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return responseData, nil
}

/*
func parseInput(input string) (data.PersonalData, error) {
	parts := strings.Split(input, "/")
	if len(parts) != 3 {
		//incorrect input format
		return data.PersonalData{}, fmt.Errorf("niepoprawny format wejściowy")
	}
	return data.PersonalData{
		FirstName:           parts[0],
		LastName:            parts[1],
		DriverLicenseNumber: parts[2],
	}, nil
}
*/
