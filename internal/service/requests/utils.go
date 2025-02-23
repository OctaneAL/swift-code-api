package requests

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RetrieveStringParam(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}
