package requests

import (
	"encoding/json"
	"net/http"
	"regexp"

	val "github.com/go-ozzo/ozzo-validation/v4"
)

type AddSwiftCodeRequest struct {
	Address         string `json:"address"`
	BankName        string `json:"bankName"`
	CountryISO2Code string `json:"countryISO2"`
	CountryName     string `json:"countryName"`
	IsHeadquarter   *bool  `json:"isHeadquarter"`
	SwiftCode       string `json:"swiftCode"`
}

func NewAddSwiftCodeRequest(r *http.Request) (req AddSwiftCodeRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, val.Errors{"body": err}
	}

	return req, val.Errors{
		"address":       val.Validate(req.Address, val.Required),
		"bankName":      val.Validate(req.BankName, val.Required),
		"countryISO2":   val.Validate(req.CountryISO2Code, val.Required, val.Length(2, 2), val.Match(regexp.MustCompile("^[A-Z]+$"))),
		"countryName":   val.Validate(req.CountryName, val.Required),
		"isHeadquarter": val.Validate(req.IsHeadquarter, val.NotNil, val.In(true, false)),
		"swiftCode":     val.Validate(req.SwiftCode, val.Required, val.Match(regexp.MustCompile("^[A-Z0-9]+$"))),
	}.Filter()
}
