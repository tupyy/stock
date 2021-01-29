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
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
	"github.com/tupyy/stock/client"
	"github.com/tupyy/stock/client/operations"
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("arguments missing")
			return
		}

		client := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{Host: "localhost:18080"})

		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			for range interruptChan {
				cancel()
			}
		}()

		// save position
		fmt.Print("\033[s")

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {

			var values []float64
			var lastValue float64
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()
			defer wg.Done()
			for {
				select {
				default:
					params := operations.NewGetStockParams()
					label := strings.ToUpper(args[1])
					params.Label = &label

					stockValues, err := client.Operations.GetStock(params)
					if err == nil {
						currValue := stockValues.Payload.Values[0].Value
						if currValue != lastValue {
							lastValue = currValue
							values = append(values, currValue)
							graph := asciigraph.Plot(values)

							fmt.Print("\033[u\033[K")
							fmt.Println(graph)
						}
					}
				case <-ctx.Done():
					return
				}
			}
		}()

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(plotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// plotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
