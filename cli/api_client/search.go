package api_client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mxmvncnt/packsearch/cli/util"
	"github.com/mxmvncnt/packsearch/server/database"
)

func Search(packageName string) []database.Package {
	response, err := http.Get(fmt.Sprintf("%s/search/%s", ApiHostname, packageName))
	if err != nil {
		util.ExitWithError(err.Error())
		return []database.Package{}
	}
	defer response.Body.Close()

	var result []database.Package

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		util.ExitWithError("Error parsing JSON response for package search")
	}

	return result
}
