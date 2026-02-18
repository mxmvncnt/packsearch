package api_client

import (
	"encoding/json"
	"fmt"
	"go/types"
	"net/http"

	"github.com/mxmvncnt/packsearch/cli/util"
	"github.com/mxmvncnt/packsearch/server/database"
)

func Search(packageName string) database.Package {
	response, err := http.Get(fmt.Sprintf("%s/search/%s", ApiHostname, packageName))
	if err != nil {
		util.ExitWithError(err.Error())
		return []types.Package{}
	}
	defer response.Body.Close()

	var result []types.Package

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		util.ExitWithError("Error parsing JSON response for package search")
	}

	return result
}
