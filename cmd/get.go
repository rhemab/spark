package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TylerBrock/colorjson"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send a GET request.",
	Long:  `Send a GET request. Example: spark get 'https://swapi.dev/api/people/1'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`Type: spark get "<url>" --header "<key>=<value>,<key>=<value>"`)
			return
		}
		url := args[0]
		fmt.Println("GET", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating new request:", err)
		}
		req.Header.Add("accept", "application/json")

		config := readConfig(false)

		// check if URL is already saved in config. If so, add saved headers to request
		for _, item := range config {
			for i := len(url); i > 10; i-- {
				if url[:i] == item.URL {
					for key, value := range item.Headers {
						req.Header.Add(key, value)
					}
				}
			}
		}

		// add headers from header flag
		for key, value := range HeaderFlag {
			req.Header.Add(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error making get request:", err)
		}
		if res.StatusCode != 200 {
			fmt.Println("Error:", res.Status)
			return
		}

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Could not read response body: %s\n", err)
		}
		var obj map[string]interface{}
		json.Unmarshal(resBody, &obj)
		colorJSON(obj)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func colorJSON(v any) {
	f := colorjson.NewFormatter()
	f.Indent = 4
	s, _ := f.Marshal(v)
	fmt.Println(string(s))
}
