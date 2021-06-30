/*
Copyright © 2021 Agrim Asthana <dev@agrimasthana.com>

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
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "hashmatch",
	Short: "print hashes of files in tabular form",
	Long: `A cli tool that will print the hashes of different files sent as arguments
.`,
	Run: func(cmd *cobra.Command, args []string) {
		exitcode := 0
		data := [][]string{}
		for _, v := range args {
			f, err := os.Open(v)
			if err != nil {
				log.Fatal(err)
				os.Exit(exitcode - 1)
			}
			defer f.Close()
			h := md5.New()
			if _, err := io.Copy(h, f); err != nil {
				os.Exit(exitcode - 1)
			}
			data = append(data, []string{f.Name(), fmt.Sprintf("%x", h.Sum(nil))})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"File", "md5sum"})
		for _, v := range data {
			table.Append(v)
		}

		table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})

		if len(data) >= 2 {
			if strings.Compare(data[0][1], data[1][1]) == 0 {
				table.SetFooter([]string{"Files match ✅", ""})
			} else {
				table.SetFooter([]string{"Files don't match ❌", ""})
			}
		}
		table.Render()
		os.Exit(exitcode)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
