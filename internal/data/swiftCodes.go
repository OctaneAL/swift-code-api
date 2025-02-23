package data

type SwiftCodesQ interface {
	New() SwiftCodesQ

	Upsert(swiftCode ...SwiftCode) error
	Get() (*SwiftCode, error)
	Select() ([]SwiftCode, error)
	Delete() error

	FilterBySwiftCode(swiftCode string) SwiftCodesQ
	FilterByHeadquarter(association string) SwiftCodesQ
}

type SwiftCode struct {
	ID             string `db:"id" structs:"id"`
	CountryISOCode string `db:"country_iso_code" structs:"country_iso_code"`
	SwiftCode      string `db:"swift_code" structs:"swift_code"`
	CodeType       string `db:"code_type" structs:"code_type"`
	Name           string `db:"name" structs:"name"`
	Address        string `db:"address" structs:"address"`
	TownName       string `db:"town_name" structs:"town_name"`
	CountryName    string `db:"country_name" structs:"country_name"`
	TimeZone       string `db:"time_zone" structs:"time_zone"`
}
