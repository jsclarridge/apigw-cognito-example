package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "apigw-cognito-example",
	Run: func(cmd *cobra.Command, args []string) {
		config := &clientcredentials.Config{
			ClientID:     viper.GetString("clientID"),
			ClientSecret: viper.GetString("clientSecret"),
			TokenURL:     viper.GetString("tokenEndpoint"),
			Scopes:       []string{viper.GetString("scope")},
		}
		ctx := context.Background()
		client := config.Client(ctx)
		resp, err := client.Get(viper.GetString("appURL"))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", body)
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
