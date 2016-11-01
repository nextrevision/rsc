package helper

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func WriteTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeaderLine(false)
	table.SetRowLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnSeparator("\t")
	table.SetHeader(header)
	table.AppendBulk(rows)
	table.Render()
}
