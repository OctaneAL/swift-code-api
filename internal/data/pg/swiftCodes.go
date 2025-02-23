package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/OctaneAL/swift-code-api/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

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
