package cmd

import (
	"fmt"

	"github.com/golang-boilerplate/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "app-example",
	Short: "Short Description of your application",
	Long:  "Longer Descriptition of your application",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(config.LoadConfig)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(fmt.Sprintf("%#v", err))
	}

	viper.AutomaticEnv()

	viper.ReadInConfig()
}
