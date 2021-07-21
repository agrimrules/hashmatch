package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/thoas/go-funk"
)

// CreateTable returns a formatted table to render
func CreateTable(rows []HashResults, algo string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File", algo})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})

	for _, v := range rows {
		table.Append([]string{v.Filename, v.Hash})
	}
	if len(rows) >= 2 {
		if strings.Compare(rows[0].Hash, rows[1].Hash) == 0 {
			table.SetFooter([]string{"Files match ✅", ""})
		} else {
			table.SetFooter([]string{"Files don't match ❌", ""})
		}
	}
	return table
}

//CreateDirTable returns a two column table to render
func CreateDirTable(rows []HashResults, algo string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Source File", algo, "Destination File", algo})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})
	if len(rows) >= 1 {
		table.SetFooter([]string{"Files don't match ❌", "", "", ""})
		dualRows := funk.Chunk(rows, 2).([][]HashResults)
		for _, v := range dualRows {
			table.Append([]string{v[0].Filename, v[0].Hash, v[1].Filename, v[1].Hash})
		}
	}
	return table
}

//HandleOutput is used to provide the output in the required format
func HandleOutput(op []HashResults, h string, format string, isDir bool, clearFooter bool) {
	var table *tablewriter.Table
	switch format {
	case "table":
		if !isDir {
			table = CreateTable(op, h)
		} else {
			table = CreateDirTable(op, h)
		}
		if clearFooter {
			table.ClearFooter()
		}
		table.Render()
		os.Exit(0)
	case "json":
		outputArray := JSONHashResults{Results: op, Algo: h}
		if clearFooter {
			s, _ := json.MarshalIndent(struct {
				Results []HashResults `json:"results"`
				Algo    string        `json:"algo"`
			}{Results: op, Algo: h}, "", " ")
			fmt.Println(string(s))
			os.Exit(0)
		} else {
			if len(op) >= 2 {
				if strings.Compare(op[0].Hash, op[1].Hash) == 0 {
					outputArray.Match = true
				} else {
					outputArray.Match = false
				}
			}
		}
		s, _ := json.MarshalIndent(outputArray, "", " ")
		fmt.Println(string(s))
		os.Exit(0)
	}
}
