package responses

import "github.com/OctaneAL/swift-code-api/internal/data"

type CountrySwiftCodes struct {
	CountryISO2 string             `json:"countryISO2"`
	CountryName string             `json:"countryName"`
	SwiftCodes  []SwiftCodeDetails `json:"swiftCodes"`
}

func NewCountrySwiftCodes(swiftCodes []data.SwiftCode) CountrySwiftCodes {
	response := CountrySwiftCodes{}

	for _, swiftCode := range swiftCodes {
		response.CountryISO2 = swiftCode.CountryISOCode
		response.CountryName = swiftCode.CountryName
		response.SwiftCodes = append(response.SwiftCodes, NewSwiftCodeDetails(swiftCode))
	}

	return response
}

func NewSwiftCodeDetails(swiftCode data.SwiftCode) SwiftCodeDetails {
	return SwiftCodeDetails{
		Address:       swiftCode.Address,
		BankName:      swiftCode.Name,
		CountryISO2:   swiftCode.CountryISOCode,
		CountryName:   swiftCode.CountryName,
		IsHeadquarter: len(swiftCode.SwiftCode) >= 3 && swiftCode.SwiftCode[len(swiftCode.SwiftCode)-3:] == "XXX",
		SwiftCode:     swiftCode.SwiftCode,
	}
}
