package util

import (
	"fmt"
	"os"
)

func ExitWithError(message string) {
	fmt.Println(message)
	os.Exit(1)
}
