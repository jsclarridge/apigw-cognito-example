package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jsclarridge/amazon-cognito-client/internal/api"
	"github.com/jsclarridge/amazon-cognito-client/internal/auth"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "amazon-cognito-client",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := auth.GetAccessToken(
			viper.GetString("tokenEndpoint"),
			viper.GetString("clientID"),
			viper.GetString("clientSecret"),
			"",
		)
		if err != nil {
			log.Fatal(err)
		}

		coffee, err := api.CallAPI(
			viper.GetString("appURL"),
			accessToken,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(coffee)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.BindEnv("clientID", "CLIENT_ID"); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindEnv("clientSecret", "CLIENT_SECRET"); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindEnv("tokenEndpoint", "TOKEN_ENDPOINT"); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindEnv("scope", "OAUTH_SCOPE"); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindEnv("appURL", "APP_URL"); err != nil {
		log.Fatal(err)
	}
	rootCmd.Flags().String("clientID", "", "Client ID")
	rootCmd.Flags().String("clientSecret", "", "Client Secret")
	rootCmd.Flags().String("tokenEndpoint", "", "OAuth2 token endpoint")
	rootCmd.Flags().String("scope", "", "OAuth2 scope")
	rootCmd.Flags().String("appURL", "", "Application URL")
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}
