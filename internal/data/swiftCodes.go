package data

type SwiftCodesQ interface {
	New() SwiftCodesQ

	Upsert(swiftCode ...SwiftCode) error
	Get() (*SwiftCode, error)
	Select() ([]SwiftCode, error)
	Delete() error

	FilterBySwiftCode(swiftCode string) SwiftCodesQ
	FilterByHeadquarter(association string) SwiftCodesQ
	FilterByCountryISO2Code(countryISO2Code string) SwiftCodesQ
}

type SwiftCode struct {
	ID              string `db:"id" structs:"id"`
	CountryISO2Code string `db:"country_iso2_code" structs:"country_iso2_code"`
	SwiftCode       string `db:"swift_code" structs:"swift_code"`
	CodeType        string `db:"code_type" structs:"code_type"`
	BankName        string `db:"bank_name" structs:"bank_name"`
	Address         string `db:"address" structs:"address"`
	TownName        string `db:"town_name" structs:"town_name"`
	CountryName     string `db:"country_name" structs:"country_name"`
	TimeZone        string `db:"time_zone" structs:"time_zone"`
}
