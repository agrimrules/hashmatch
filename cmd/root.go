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
	"encoding/json"
	"fmt"
	"os"

	"go.agrim.dev/hashmatch/utils"
	"github.com/spf13/cobra"
)

var version = "dev"
var v bool
var h string
var o string

var rootCmd = &cobra.Command{
	Use:   "hashmatch",
	Short: "print hashes of files in tabular form",
	Long: `A cli tool that will print the hashes of different files sent as arguments
.`,
	Run: func(cmd *cobra.Command, args []string) {
		exitcode := 0
		isArg0Directory := false
		isArg1Directory := false

		if len(args) >= 1 && args[0] != "" {
			isArg0Directory = utils.IsDirectory(args[0])
		}
		if len(args) >= 2 && args[1] != "" {
			isArg1Directory = utils.IsDirectory(args[1])
		}
		if v {
			fmt.Println(version)
			os.Exit(exitcode)
		}
		switch len(args) {
		case 1:
			var d []utils.HashResults
			if isArg0Directory {
				d = utils.GetHashesForFiles(utils.ReturnFilesInFolder(args[0]), h)
			} else {
				d = utils.GetHashesForFiles(args, h)
			}
			utils.HandleOutput(d, h, o, false, true)

		case 2:
			if isArg0Directory && isArg1Directory {
				d1 := utils.GetHashesForFiles(utils.ReturnFilesInFolder(args[0]), h)
				d2 := utils.GetHashesForFiles(utils.ReturnFilesInFolder(args[1]), h)
				areEqual, diff := utils.HashesAreEqual(d1, d2)
				if !areEqual {
					utils.HandleOutput(diff, h, o, true, false)
				}
				if o == "table" {
					fmt.Println("Files match ✅")
				}
				if o == "json" {
					op, _ := json.MarshalIndent(struct {
						Matched bool `json:"matched"`
					}{Matched: true}, "", " ")
					fmt.Println(string(op))
				}
				os.Exit(exitcode)
			} else if !isArg0Directory && !isArg1Directory {
				data := utils.GetHashesForFiles(args, h)
				utils.HandleOutput(data, h, o, false, false)
			}
			fmt.Println("Can only compare two folders or two files")

		default:
			fmt.Println("Invalid Arguments")
			os.Exit(-1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().BoolVarP(&v, "version", "v", false, "Print version")
	rootCmd.Flags().StringVarP(&h, "hash", "", "md5sum", "Specify hash algorithm to use")
	rootCmd.Flags().StringVarP(&o, "output", "o", "table", "Specify the output format to use")
	cobra.CheckErr(rootCmd.Execute())
}
