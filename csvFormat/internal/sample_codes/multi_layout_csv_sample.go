package samplecodes

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/gocarina/gocsv"
)

type Summary struct {
	RecordType string
	Summary    string
}

type Book struct {
	RecordType string
	Title      string
	Page       uint16
	Price      uint16
	Author     string
}

type singleCSVReader struct {
	record []string
}

func (r singleCSVReader) Read() ([]string, error) {
	return r.record, nil
}

func (r singleCSVReader) ReadAll() ([][]string, error) {
	return [][]string{r.record}, nil
}

func ReadMultiLayoutCSVSample() {
	s := `summary,3件
book,サンプル本1,100,1500,sample_author1
book,サンプル本2,200,1200,sample_author2
book,サンプル本3,120,1000,sample_author1`

	r := csv.NewReader(strings.NewReader(s)) // 標準パッケージを使用して読み込み
	r.FieldsPerRecord = -1                   // レコード数が異なるため、チェックを無効化
	all, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range all {
		// RecordTypeによって読み込み先を分岐する
		if record[0] == "summary" {
			var summaries []Summary
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record: record}, &summaries); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("summary行の読み込み: %+v\n", summaries[0])
		} else {
			var books []Book
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record: record}, &books); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("book行の読み込み： %+v\n", books[0])
		}
	}
}
