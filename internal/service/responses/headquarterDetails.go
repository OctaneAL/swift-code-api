package responses

type HeadquarterDetails struct {
	SwiftCodeDetails
	Branches []SwiftCodeDetails `json:"branches"`
}
