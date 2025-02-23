package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/OctaneAL/swift-code-api/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

// TODO: provide better namings for columns
const (
	swiftCodesTableName            = "swift_codes"
	swiftCodesIDColumn             = "id"
	swiftCodesCountryISOCodeColumn = "country_iso_code"
	swiftCodesSwiftCodeColumn      = "swift_code"
	swiftCodesCodeTypeColumn       = "code_type"
	swiftCodesNameColumn           = "name"
	swiftCodesAddressColumn        = "address"
	swiftCodesTownNameColumn       = "town_name"
	swiftCodesCountryNameColumn    = "country_name"
	swiftCodesTimeZoneColumn       = "time_zone"
)

var swiftCodesColumns = []string{
	swiftCodesIDColumn,
	swiftCodesCountryISOCodeColumn,
	swiftCodesSwiftCodeColumn,
	swiftCodesCodeTypeColumn,
	swiftCodesNameColumn,
	swiftCodesAddressColumn,
	swiftCodesTownNameColumn,
	swiftCodesCountryNameColumn,
	swiftCodesTimeZoneColumn,
}

type swiftCodesQ struct {
	db       *pgdb.DB
	selector sq.SelectBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
}

func NewSwiftCodesQ(db *pgdb.DB) data.SwiftCodesQ {
	return &swiftCodesQ{
		db:       db,
		selector: sq.Select("*").From(swiftCodesTableName),
		updater:  sq.Update(swiftCodesTableName),
		deleter:  sq.Delete(swiftCodesTableName),
	}
}

func (q swiftCodesQ) New() data.SwiftCodesQ {
	return NewSwiftCodesQ(q.db.Clone())
}

func (q swiftCodesQ) Upsert(swiftCodes ...data.SwiftCode) error {
	if len(swiftCodes) == 0 {
		return nil
	}

	query := sq.Insert(swiftCodesTableName).Columns(swiftCodesColumns...)
	for _, swiftCode := range swiftCodes {
		query = query.Values(structs.Values(swiftCode)...)
	}
	// TODO: implement this via library, not by hand
	query = query.Suffix("ON CONFLICT (swift_code) DO UPDATE SET country_iso_code = EXCLUDED.country_iso_code, code_type = EXCLUDED.code_type, name = EXCLUDED.name, address = EXCLUDED.address, town_name = EXCLUDED.town_name, country_name = EXCLUDED.country_name, time_zone = EXCLUDED.time_zone")

	return q.db.Exec(query)
}

func (q swiftCodesQ) FilterBySwiftCode(swiftCode string) data.SwiftCodesQ {
	return q.withFilters(sq.Eq{swiftCodesSwiftCodeColumn: swiftCode})
}

func (q swiftCodesQ) FilterByHeadquarter(headquarterCode string) data.SwiftCodesQ {
	return q.withFilters(sq.Expr("LEFT("+swiftCodesSwiftCodeColumn+", 8) = ?", headquarterCode))
}

func (q swiftCodesQ) FilterByCountryISO2Code(countryISO2Code string) data.SwiftCodesQ {
	return q.withFilters(sq.Eq{swiftCodesCountryISOCodeColumn: countryISO2Code})
}

func (q swiftCodesQ) withFilters(stmt interface{}) data.SwiftCodesQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)
	q.deleter = q.deleter.Where(stmt)

	return q
}

func (q swiftCodesQ) Get() (*data.SwiftCode, error) {
	var result data.SwiftCode
	err := q.db.Get(&result, q.selector)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q swiftCodesQ) Select() ([]data.SwiftCode, error) {
	var result []data.SwiftCode

	return result, q.db.Select(&result, q.selector)
}

func (q swiftCodesQ) Delete() error {
	return q.db.Exec(q.deleter)
}
