package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	dbAdapter "github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/adapters/db"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = dbAdapter.NewProductDB(db)
var productService = application.ProductService{Persistence: productDb}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "products-application",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.products-application.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".products-application")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
