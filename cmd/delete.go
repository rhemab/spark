package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Send a DELETE request.",
	Long:  `Send a DELETE request. Example: spark delete 'https://swapi.dev/api/people/1'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Add url as argument.")
			return
		}
		url := args[0]
		req, err := http.NewRequest("Delete", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		for key, value := range HeaderFlag {
			req.Header.Add(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != 200 {
			fmt.Println(res.Status)
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
	rootCmd.AddCommand(deleteCmd)
}
