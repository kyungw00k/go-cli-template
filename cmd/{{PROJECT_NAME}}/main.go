package main

import (
	"fmt"
	"os"

	"{{MODULE_PATH}}/api"
	"{{MODULE_PATH}}/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		if apiErr, ok := api.AsAPIError(err); ok {
			fmt.Fprintln(os.Stderr, apiErr.Error())
			os.Exit(apiErr.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
