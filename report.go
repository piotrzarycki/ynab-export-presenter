package main

import (
	"bufio"
	"encoding/csv"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
)

func main() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetAutoMergeCells(true)
	table.SetBorder(false)
	f, err := os.Open("report.csv")

	if err != nil {
	}

	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'
	i := 0
	for {
		record, err := r.Read()
		if len(record) == 0 {
			break
		}

		if i == 0 {
			table.SetHeader(record)
		} else {

			table.Append(record)
		}
		if err == io.EOF {
			break
		}

		i++
	}

	if err != nil {
	}

	table.Render()
}
