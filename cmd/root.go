package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spark",
	Short: "Spark is a CLI tool for testing http requests.",
	Long:  `Spark is a CLI tool for testing http requests.`,
}

var HeaderFlag map[string]string
var cfgFile string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringToStringVar(&HeaderFlag, "header", HeaderFlag, `Set request header. Example: --header "accept=application/json,secret-key=lksdfhksj,key=value"`)
}
