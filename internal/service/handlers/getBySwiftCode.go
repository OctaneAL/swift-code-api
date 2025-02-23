package handlers

import (
	"errors"
	"net/http"

	"github.com/OctaneAL/swift-code-api/internal/service/requests"
	"github.com/OctaneAL/swift-code-api/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBySwiftCode(w http.ResponseWriter, r *http.Request) {
	swiftCode := requests.RetrieveStringParam(r, "swiftCode")

	// Check for headquarter
	// TODO: Return message if swift code is not found
	if len(swiftCode) >= 3 && swiftCode[len(swiftCode)-3:] == "XXX" {
		response, err := getHeadquarterDetails(swiftCode, r)
		if err != nil {
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.Render(w, response)
		return
	} else {
		response, err := getBranchDetails(swiftCode, r)
		if err != nil {
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.Render(w, response)
		return
	}
}

func getHeadquarterDetails(swiftCode string, r *http.Request) (*responses.HeadquarterDetails, error) {
	headquarterRecord, err := SwiftCodesQ(r).FilterBySwiftCode(swiftCode).Get()
	if err != nil {
		Log(r).WithError(err).WithField("swiftCode", swiftCode).Error("failed to get swift code")
		return nil, err
	}
	if headquarterRecord == nil {
		Log(r).WithField("swiftCode", swiftCode).Error("no swift code found")
		return nil, errors.New("no swift code found")
	}

	headquarter := responses.HeadquarterDetails{
		SwiftCodeDetails: responses.SwiftCodeDetails{
			Address:       headquarterRecord.Address,
			BankName:      headquarterRecord.Name,
			CountryISO2:   headquarterRecord.CountryISOCode,
			CountryName:   headquarterRecord.CountryName,
			IsHeadquarter: true,
			SwiftCode:     headquarterRecord.SwiftCode,
		},
		Branches: []responses.SwiftCodeDetails{},
	}

	branches, err := SwiftCodesQ(r).FilterByHeadquarter(swiftCode[:len(swiftCode)-3]).Select()
	if err != nil {
		Log(r).WithError(err).WithField("swiftCode", swiftCode).Error("failed to get branches")
		return nil, err
	}

	for _, branchRecord := range branches {
		if len(branchRecord.SwiftCode) >= 3 && branchRecord.SwiftCode[len(branchRecord.SwiftCode)-3:] == "XXX" {
			continue
		}

		headquarter.Branches = append(headquarter.Branches, responses.SwiftCodeDetails{
			Address:       branchRecord.Address,
			BankName:      branchRecord.Name,
			CountryISO2:   branchRecord.CountryISOCode,
			CountryName:   branchRecord.CountryName,
			SwiftCode:     branchRecord.SwiftCode,
			IsHeadquarter: false,
		})
	}

	return &headquarter, nil
}

func getBranchDetails(swiftCode string, r *http.Request) (*responses.SwiftCodeDetails, error) {
	branchRecord, err := SwiftCodesQ(r).FilterBySwiftCode(swiftCode).Get()
	if err != nil {
		Log(r).WithError(err).WithField("swiftCode", swiftCode).Error("failed to get swift code")
		return nil, err
	}
	if branchRecord == nil {
		Log(r).WithField("swiftCode", swiftCode).Error("no swift code found")
		return nil, errors.New("no swift code found")
	}

	branch := responses.SwiftCodeDetails{
		Address:       branchRecord.Address,
		BankName:      branchRecord.Name,
		CountryISO2:   branchRecord.CountryISOCode,
		CountryName:   branchRecord.CountryName,
		IsHeadquarter: false,
		SwiftCode:     branchRecord.SwiftCode,
	}

	return &branch, nil
}
