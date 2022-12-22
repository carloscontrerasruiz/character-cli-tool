/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var pokemonName string

// pokemonCmd represents the pokemon command
var PokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Find your pokemon",
	Long: `Find your pokemon`,
	Run: func(cmd *cobra.Command, args []string) {
		callPokemonApi()
	},
}

func init() {
	PokemonCmd.Flags().StringVarP(&pokemonName, "name", "n", "charizard", "The name of the pokemon")
	PokemonCmd.MarkFlagRequired("name")
}

func callPokemonApi(){
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
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
	var pokemonResp Pokedex
	if err := json.NewDecoder(res.Body).Decode(&pokemonResp); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Printf("Name of the pokemon: %s\n", pokemonResp.Name)
	fmt.Printf("Number of the pokemon: %d\n", pokemonResp.ID)
	fmt.Printf("Weihgt of the pokemon: %d\n", pokemonResp.Weight)
}