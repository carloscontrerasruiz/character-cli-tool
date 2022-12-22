/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package rick

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rickCmd represents the rick command
var RickCmd = &cobra.Command{
	Use:   "rick",
	Short: "Find rick and morty character",
	Long: `Find rick and morty character`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		// Run the custom validation logic
		if len(args[0]) >= 4 {
		  return nil
		}

		return fmt.Errorf("The name of the charcter must have more tha 3 characters")
	  },
	Run: func(cmd *cobra.Command, args []string) {
		callRickApi(args[0])
	},
}

func init() {
	//rootCmd.AddCommand(rickCmd)
}

func callRickApi(name string){
	url := "https://rickandmortyapi.com/api/character/?name=" + name
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err.Error())
		os.Exit(1)
	}

	if res.StatusCode >= 300{
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	defer res.Body.Close()
	var apiResponse ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	character := apiResponse.Results[0]
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Println("=========================")
	fmt.Printf("Number of characters find with that name: %d\n", apiResponse.Info.Count)
	fmt.Println("Number of characters shown: 1")
	fmt.Println("=========================")
	fmt.Printf("Name of the Rick and Morty character: %s\n", character.Name)
	fmt.Printf("Species: %s\n", character.Species)
	fmt.Printf("Status: %s\n", character.Status)
	fmt.Printf("ID: %d\n", character.ID)
}
