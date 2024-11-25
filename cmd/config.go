package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type UrlConfig struct {
	URL     string
	Headers map[string]string
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set default headers for base URLs.",
	Long:  `Set default headers for base URLs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Add a base URL with the header flag to the config command: spark config '<baseURL>' --header '<key>=<value>,<key>=<value>'")
			readConfig(true)
			return
		}
		currentConfig := readConfig(false)
		url := args[0]
		newConfig := []UrlConfig{
			{
				URL:     url,
				Headers: HeaderFlag,
			},
		}

		// check if URL is already saved in config.
		for _, item := range currentConfig {
			for i := len(url); i > 10; i-- {
				if url[:i] == item.URL {
					var replaceURL string
					fmt.Println("URL is already saved, would you like to replace it (y/n)?")
					fmt.Scan(&replaceURL)
					if replaceURL != "y" {
						return
					}
					fmt.Println("Updating URL config...")
					item.URL = url
					item.Headers = HeaderFlag
					writeConfig(currentConfig)
					fmt.Println("Config saved.")
					return
				}
			}
		}

		if currentConfig == nil {
			fmt.Println("No config found, creating config file.")
			writeConfig(newConfig)
		} else {
			writeConfig(append(currentConfig, newConfig...))
		}
		fmt.Println("Config saved.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func readConfig(showConfig bool) (currentConfig []UrlConfig) {
	configBytes, err := os.ReadFile("sparkConfig.json")
	if err != nil {
		return nil
	}
	json.Unmarshal(configBytes, &currentConfig)

	if showConfig {
		var config interface{}
		json.Unmarshal(configBytes, &config)
		fmt.Println("Current Config:")
		colorJSON(config)
	}
	return
}

func writeConfig(v []UrlConfig) {
	file, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatal("Error json.Marshal", err)
	}
	os.WriteFile("sparkConfig.json", file, 0664)
}
