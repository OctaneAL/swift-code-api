package parser

import (
	"encoding/csv"
	"os"

	"github.com/OctaneAL/swift-code-api/internal/config"
	"github.com/OctaneAL/swift-code-api/internal/data"
	"github.com/OctaneAL/swift-code-api/internal/data/pg"
	"github.com/google/uuid"
)

func Run(cfg config.Config) {
	if err := processData(cfg); err != nil {
		panic(err)
	}
}

func processData(cfg config.Config) error {
	cfg.Log().Info("start processing data")

	file, err := os.Open(cfg.DataPath().DataPath)
	if err != nil {
		cfg.Log().WithError(err).Error("failed to open file")
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var swiftCodes []data.SwiftCode
	for i, row := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		if len(row) != 8 {
			cfg.Log().Errorf("invalid row: %v", row)
			continue
		}

		// TODO: is there a better way to Unpack the row into a struct?
		swiftCode := data.SwiftCode{
			ID:              uuid.New().String(),
			CountryISO2Code: row[0],
			SwiftCode:       row[1],
			CodeType:        row[2],
			BankName:        row[3],
			Address:         row[4],
			TownName:        row[5],
			CountryName:     row[6],
			TimeZone:        row[7],
		}

		swiftCodes = append(swiftCodes, swiftCode)
	}

	cfg.Log().Info("finish processing data")

	swiftCodesQ := pg.NewSwiftCodesQ(cfg.DB().Clone())
	if err := swiftCodesQ.Upsert(swiftCodes...); err != nil {
		cfg.Log().WithError(err).Error("failed to upsert swift codes")
		return err
	}

	return nil
}
