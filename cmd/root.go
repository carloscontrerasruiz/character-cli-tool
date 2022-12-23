package cmd

import (
	"os"

	"github.com/carloscontrerasruiz/character/cmd/pokemon"
	"github.com/carloscontrerasruiz/character/cmd/rick"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "character",
	Short: "A brief description of your application",
	Long:  `A longer description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enabling verbose mode")

	addSubcommands()

}

func addSubcommands() {
	rootCmd.AddCommand(pokemon.PokemonCmd)
	rootCmd.AddCommand(rick.RickCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	//Adding default config path
	//default is $HOME/character.yaml
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	//Setting default values
	viper.SetDefault("verboseSpeech", "This is the default line that should be apper in verbose mode")

	// Search config in home directory with name ".character" (without extension).
	viper.SetConfigName("character")
	viper.SetConfigType("json")
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil {
		cobra.CheckErr(err)
	}

}
