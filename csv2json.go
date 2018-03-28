package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("program needs the csv file")
	}
	csvFileName := os.Args[1]

	csvFile, err := ioutil.ReadFile(csvFileName)
	if err != nil {
		log.Fatalf("open config file failed: %s, %s", csvFileName, err)
	}

	rows := strings.Split(string(csvFile), "\n")
	var headers string
	if len(rows) > 0 {
		headers = rows[0]
		rows = rows[1:]
	}
	headerColumns := strings.Split(headers, ",")

	items := []map[string]string{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		dataColumns := strings.Split(row, ",")
		if len(headerColumns) != len(dataColumns) {
			panic(fmt.Sprintf("invalid csv: %s,%s", headers, row))
		}
		datum := make(map[string]string)
		for i := range headerColumns {
			datum[headerColumns[i]] = dataColumns[i]
		}
		items = append(items, datum)
	}

	data, _ := json.Marshal(items)

	fmt.Print(string(data))
}
