package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/OctaneAL/swift-code-api/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteSwiftCode(w http.ResponseWriter, r *http.Request) {
	swiftCode := requests.RetrieveStringParam(r, "swiftCode")

	err := SwiftCodesQ(r).FilterBySwiftCode(swiftCode).Delete()
	if err != nil {
		Log(r).WithError(err).Error("failed to delete swift code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := map[string]string{"message": "Swift code deleted successfully"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		Log(r).WithError(err).Error("failed to write response")
		ape.RenderErr(w, problems.InternalError())
	}
}
