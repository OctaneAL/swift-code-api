package handlers

import (
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

	w.WriteHeader(http.StatusNoContent)
}
