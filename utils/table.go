package utils

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/thoas/go-funk"
)

func CreateTable(rows []HashResults) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File", "md5sum"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})

	for _, v := range rows {
		table.Append([]string{v.filename, v.hash})
	}
	if len(rows) >= 2 {
		if strings.Compare(rows[0].hash, rows[1].hash) == 0 {
			table.SetFooter([]string{"Files match ✅", ""})
		} else {
			table.SetFooter([]string{"Files don't match ❌", ""})
		}
	}
	return table
}

func CreateDirTable(rows []HashResults) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Source File", "md5sum", "Destination File", "md5sum"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})
	if len(rows) >= 1 {
		table.SetFooter([]string{"Files don't match ❌", "", "", ""})
		dualRows := funk.Chunk(rows, 2).([][]HashResults)
		for _, v := range dualRows {
			table.Append([]string{v[0].filename, v[0].hash, v[1].filename, v[1].hash})
		}
	}
	return table
}
