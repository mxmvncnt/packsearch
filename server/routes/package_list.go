package routes

import (
	"net/http"

	"github.com/mxmvncnt/packsearch/utils"
)

func (handler *RoutesHandler) GetPackagesList(w http.ResponseWriter, r *http.Request) error {

	result, err := handler.db.GetAllPackages(r.Context())
	if err != nil {
		return err
	}

	utils.SendJsonResponse(w, http.StatusOK, result)

	return nil
}
