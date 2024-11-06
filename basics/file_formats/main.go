package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

var records [][]string

type User struct {
	Userid int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    string `json:"age"`
}

func main() {

	f, _ := os.Open("data.csv")

	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		records = append(records, record)
	}

	f2, _ := os.OpenFile("./data_sorted.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	w := csv.NewWriter(f2)
	defer w.Flush()
	sort.Slice(records, func(i, j int) bool {
		return records[i][1] < records[j][1]
	})
	for _, record := range records {
		w.Write(record)
	}

	data := make(map[string][]User)
	jsonFile, _ := os.ReadFile("users.json")
	json.Unmarshal(jsonFile, &data)

	fmt.Println(data["users"])
}
