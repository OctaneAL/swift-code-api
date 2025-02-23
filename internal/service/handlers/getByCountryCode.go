package handlers

import (
	"net/http"

	"github.com/OctaneAL/swift-code-api/internal/service/requests"
	"github.com/OctaneAL/swift-code-api/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := requests.RetrieveStringParam(r, "countryISO2code")

	swiftCodes, err := SwiftCodesQ(r).FilterByCountryISO2Code(countryCode).Select()
	if err != nil {
		Log(r).WithError(err).WithField("countryISO2code", countryCode).Error("failed to get swift codes")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(swiftCodes) == 0 {
		Log(r).WithField("countryISO2code", countryCode).Error("no swift codes found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, responses.NewCountrySwiftCodes(swiftCodes))
}
