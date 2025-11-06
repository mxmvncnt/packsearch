package routes

import (
	"net/http"

	"github.com/mxmvncnt/packsearch/utils"
)

func (handler *RoutesHandler) SearchPackage(w http.ResponseWriter, r *http.Request) error {
	searchTerm := r.PathValue("search_term")

	result, err := handler.db.FuzzySearch(r.Context(), searchTerm)
	if err != nil {
		return err
	}

	utils.SendJsonResponse(w, http.StatusOK, result)

	return nil
}
