package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
Improper use of tool.
go run ./csv_to_markdown_table.go input_file.csv`)
		os.Exit(1)
	}

	markdownTable, err := createMarkdownTable(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(markdownTable)
}

func createMarkdownTable(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	if len(records) == 0 {
		return "", fmt.Errorf("no data found in the file")
	}

	var markdownTable string

	// Create the header row
	markdownTable += "| "
	for i, header := range records[0] {
		markdownTable += header
		markdownTable += " |"
		if i < len(records[0])-1 {
			markdownTable += " "
		}
	}
	markdownTable += "\n"

	// Create the separator row
	markdownTable += "|"
	for _ = range records[0] {
		markdownTable += " --- |"
	}
	markdownTable += "\n"

	// Create the data rows
	for _, record := range records[1:] {
		markdownTable += "| "
		for i, value := range record {
			// Handle empty fields: you can choose to replace "" with "-" or leave it empty.
			if value == "" {
				markdownTable += " "
			} else {
				markdownTable += value
			}
			markdownTable += " |"
			if i < len(record)-1 {
				markdownTable += " "
			}
		}
		markdownTable += "\n"
	}

	return markdownTable, nil
}

