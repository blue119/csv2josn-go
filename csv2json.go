package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func csv2json(r io.Reader) ([]byte, error) {
	var first_row bool = true
	title_row := make([]string, 0)
	rows := make(map[string]map[string]string)
	csvReader := csv.NewReader(r)
	csvReader.TrimLeadingSpace = true
	var c int = 0

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if first_row == true {
			title_row = make([]string, len(record))
			copy(title_row, record)
			first_row = false
			continue
		}

		//  log.Println(record)
		row := make(map[string]string)
		for i, n := range title_row {
			row[n] = record[i]
		}
		//  rows = append(rows, row)
		rows[strconv.Itoa(c)] = row
		c++
	}

	data, err := json.MarshalIndent(&rows, "", "    ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tgo run csv2josn.go <csv>")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(-1)
	}

	var in_file_name = os.Args[1]

	in_file, err := os.Open(in_file_name)

	if err != nil {
		log.Fatalf("%s is not existed.\n", in_file_name)
		os.Exit(-1)
	}
	defer in_file.Close()

	jsonData, err := csv2json(in_file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
