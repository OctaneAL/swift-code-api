package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OctaneAL/swift-code-api/internal/data"
	"github.com/OctaneAL/swift-code-api/internal/service/requests"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AddSwiftCode(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAddSwiftCodeRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: fmt.Sprintf("Failed to create swift code: %v", err),
			Status: fmt.Sprintf("%d", http.StatusInternalServerError),
		})
		// ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err = SwiftCodesQ(r).Upsert(data.SwiftCode{
		SwiftCode:       request.SwiftCode,
		Address:         request.Address,
		BankName:        request.BankName,
		CountryISO2Code: request.CountryISO2Code,
		CountryName:     request.CountryName,
	}); err != nil {
		Log(r).WithError(err).Error("failed to create swift code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := map[string]string{"message": "Swift code created successfully"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		Log(r).WithError(err).Error("failed to write response")
		ape.RenderErr(w, problems.InternalError())
	}
}
