package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mxmvncnt/packsearch/server/utils"
	"github.com/mxmvncnt/packsearch/server/utils/apierror"
)

func (handler *RoutesHandler) GetVariations(w http.ResponseWriter, r *http.Request) error {
	packageID := r.PathValue("id")

	castedPackageID := pgtype.Numeric{}
	err := castedPackageID.Scan(packageID)
	if err != nil {
		return apierror.NewApiErrorWithError(
			http.StatusUnprocessableEntity,
			"bad_id",
			"Make sure package ID is an INTEGER",
			"Could not process package ID",
			err)
	}
	result, err := handler.db.GetVariations(r.Context(), castedPackageID)
	if err != nil {
		return err
	}

	utils.SendJsonResponse(w, http.StatusOK, result)

	return nil
}
