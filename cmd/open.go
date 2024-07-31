package cmd

import (
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the desired OTel UI in the default browser. Accepts the following arguments: 'prometheus' or 'p', 'grafana' or 'g', 'jaeger' or 'j'.",
}