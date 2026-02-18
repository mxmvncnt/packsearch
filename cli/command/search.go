package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mxmvncnt/packsearch/cli/api_client"
)

func SearchCommand() {

	programNames := os.Args[2]

	results := api_client.Search(programNames)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, err := fmt.Fprintln(w, "HUMAN NAME\tNAME\tVERSION")
	if err != nil {
		return
	}

	for _, result := range results {
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", result.HumanName, result.Name, result.LatestVersion)
		if err != nil {
			return
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}
