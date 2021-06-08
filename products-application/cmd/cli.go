package cmd

import (
	"fmt"

	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / Disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
}
