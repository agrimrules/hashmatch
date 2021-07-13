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
	"fmt"
	"os"

	"github.com/agrimrules/hashmatch/utils"
	"github.com/spf13/cobra"
)

var version = "dev"
var v bool

var rootCmd = &cobra.Command{
	Use:   "hashmatch",
	Short: "print hashes of files in tabular form",
	Long: `A cli tool that will print the hashes of different files sent as arguments
.`,
	Run: func(cmd *cobra.Command, args []string) {
		exitcode := 0
		// data := []utils.HashResults{}
		if v {
			fmt.Println(version)
			os.Exit(exitcode)
		}
		switch len(args) {
		case 2:
			if utils.IsDirectory(args[0]) && utils.IsDirectory(args[1]) {
				d1 := utils.GetMD5ForFiles(utils.ReturnFilesInFolder(args[0]))
				d2 := utils.GetMD5ForFiles(utils.ReturnFilesInFolder(args[1]))
				areEqual, diff := utils.HashesAreEqual(d1, d2)
				if !areEqual {
					table := utils.CreateDirTable(diff)
					table.Render()
					os.Exit(exitcode)
				}
				fmt.Println("Files match ✅")
				os.Exit(exitcode)
			} else if !utils.IsDirectory(args[0]) && !utils.IsDirectory(args[1]) {
				data := utils.GetMD5ForFiles(args)
				table := utils.CreateTable(data)
				table.Render()
				os.Exit(exitcode)
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
	cobra.CheckErr(rootCmd.Execute())
}
