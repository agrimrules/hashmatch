package utils

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
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
