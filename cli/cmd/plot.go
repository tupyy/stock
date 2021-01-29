/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
	"github.com/tupyy/stock/client"
	"github.com/tupyy/stock/client/operations"
)

var (
	height int
	width  int
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Plot the stock values of a company",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("arguments missing")
			return
		}

		client := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{Host: "localhost:18080"})

		params := operations.NewGetStocksCompanyParams()
		params.Company = strings.ToUpper(args[0])

		stockValues, err := client.Operations.GetStocksCompany(params)
		if err == nil {
			graph := asciigraph.Plot(stockValues.GetPayload().Values,
				asciigraph.Caption(args[0]),
				asciigraph.Height(height),
				asciigraph.Width(width),
			)
			fmt.Println(graph)
		}
	},
}

func init() {
	rootCmd.AddCommand(plotCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	plotCmd.Flags().IntVarP(&height, "height", "p", 20, "Heigh of the graph")
	plotCmd.Flags().IntVarP(&width, "width", "w", 100, "Width of the graph")
}
